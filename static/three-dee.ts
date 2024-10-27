import * as THREE from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls.js";
import { FirstPersonControls, FlyControls } from "three/examples/jsm/Addons.js";
import * as BufferGeometryUtils from "three/addons/utils/BufferGeometryUtils.js";

import GUI from "lil-gui";
import { tileMapData } from "./data";
import { MapBuildData } from "./interface";

interface ThreeDeeObject {
  geometry: THREE.BufferGeometry;
  material: THREE.Material;
  position?: THREE.Vector3;
  rotation?: THREE.Vector3;
}

class ThreeDeeScene {
  private scene: THREE.Scene;
  private camera: THREE.PerspectiveCamera;
  private renderer: THREE.WebGLRenderer;
  private displayElement: HTMLElement;
  private controls: FirstPersonControls;
  private directionalLight: THREE.DirectionalLight;
  private gui: GUI;
  private segments: number;

  constructor(private displayId: string, segments: number) {
    this.scene = new THREE.Scene();
    this.scene.background = null;
    this.segments = segments;
    const display = document.getElementById(this.displayId);

    if (!display) {
      throw new Error(`No element with id "${this.displayId}" found`);
    }
    this.displayElement = display;

    const container = display.parentElement;
    if (!container) {
      throw new Error(`Could not find parent element for ${this.displayId}`);
    }

    let minSize = Math.min(container.offsetWidth, container.offsetHeight);

    // Ensure the size is divisible by the number of segments and is an integer
    if (segments > 0) {
      minSize = Math.floor(minSize / this.segments) * this.segments;
    }

    this.setupRenderer(minSize, minSize);
    this.setupCamera(minSize, minSize);
    this.addLights();
    this.setupControls("");
    this.addHelpers();
    // this.setupGUI();
    this.handleWindowResize();
  }

  private setupRenderer(width: number, height: number): void {
    this.renderer = new THREE.WebGLRenderer({
      canvas: this.displayElement,
      antialias: true,
      alpha: true,
    });
    this.renderer.setSize(width, height);
  }

  private setupCamera(width: number, height: number): void {
    const halfSegment = this.segments / 2;
    this.camera = new THREE.PerspectiveCamera(75, width / height, 0.1, 1000);

    // Set the camera position directly above the scene at some height (100)
    this.camera.position.set(
      this.segments * 1.2,
      this.segments * 0.8,
      this.segments * 1.2
    );

    // Make the camera look towards the middle of the tiles at a 45 degree angle
    this.camera.lookAt(halfSegment, 0, halfSegment);
  }

  private setupControls(mode: string): void {
    switch (mode) {
      case "orbit":
        this.controls = new OrbitControls(
          this.camera,
          this.renderer.domElement
        );
        this.controls.enableDamping = true;
        this.controls.dampingFactor = 0.05;
        this.controls.screenSpacePanning = false;
        this.controls.maxPolarAngle = Math.PI / 2;
        break;
      case "firstPerson":
        this.controls = new FirstPersonControls(
          this.camera,
          this.renderer.domElement
        );
        break;
      case "fly":
        this.controls = new FlyControls(this.camera, this.renderer.domElement);
        break;
      default:
        this.controls = null;
        break;
    }
  }

  private addLights(): void {
    this.directionalLight = new THREE.DirectionalLight(0xffffff, 3);
    this.directionalLight.position.set(10, 10, 5);
    this.directionalLight.castShadow = true;

    this.scene.add(this.directionalLight);

    const light = new THREE.HemisphereLight(0xffffbb, 0x080820, 1);
    this.scene.add(light);
  }

  private addHelpers(): void {
    const axesHelper = new THREE.AxesHelper(5);
    this.scene.add(axesHelper);

    const size = 1 * this.segments;
    const divisions = this.segments;
    const gridHelper = new THREE.GridHelper(size, divisions);
    gridHelper.position.y = 12;
    gridHelper.position.x = size / 2;
    gridHelper.position.z = size / 2;
    // this.scene.add(gridHelper);
  }

  private handleWindowResize(): void {
    window.addEventListener("resize", () => {
      const { clientWidth, clientHeight } = this.displayElement;
      this.renderer.setSize(clientWidth, clientHeight);
      this.camera.aspect = clientWidth / clientHeight;
      this.camera.updateProjectionMatrix();
    });
  }

  private setupGUI(): void {
    this.gui = new GUI();
    this.gui.add(document, "title");

    // Add basic parameters to the GUI
    const cameraFolder = this.gui;
    cameraFolder.add(this.camera.position, "x", -100, 100).name("X Position");
    cameraFolder.add(this.camera.position, "y", -100, 100).name("Y Position");
    cameraFolder.add(this.camera.position, "z", -100, 100).name("Z Position");
    cameraFolder.open();

    // Add controls for the camera rotation
    const rotationFolder = this.gui.addFolder("Camera Rotation");
    rotationFolder.add(this.camera.rotation, "x", -Math.PI, Math.PI);
    rotationFolder.add(this.camera.rotation, "y", -Math.PI, Math.PI);
    rotationFolder.add(this.camera.rotation, "z", -Math.PI, Math.PI);
  }

  public addObject({
    geometry,
    material,
    position,
    rotation,
  }: ThreeDeeObject): THREE.Mesh {
    const mesh = new THREE.Mesh(geometry, material);

    if (position) {
      mesh.position.copy(position);
    }

    if (rotation) {
      mesh.rotation.set(rotation.x, rotation.y, rotation.z);
    }

    this.scene.add(mesh);

    return mesh;
  }

  public showMeshGui(mesh: THREE.Mesh): void {
    const meshFolder = this.gui.addFolder("Mesh");
    // Add controls for position
    meshFolder.add(mesh.position, "x").name("X Position");
    meshFolder.add(mesh.position, "y").name("Y Position");
    meshFolder.add(mesh.position, "z").name("Z Position");

    // Add controls for geometry dimensions
    const geometryFolder = this.gui.addFolder("Geometry");
    mesh.geometry.computeBoundingBox();
    if (mesh.geometry.boundingBox) {
      const Vector3 = new THREE.Vector3();
      mesh.geometry.boundingBox.getSize(Vector3);
      geometryFolder
        .add(Vector3, "x")
        .name("Width")
        .onChange((value) => {
          mesh.geometry.dispose();
          mesh.geometry = new THREE.BoxGeometry(value, Vector3.y, Vector3.z);
        });
      geometryFolder
        .add(Vector3, "y")
        .name("Height")
        .onChange((value) => {
          mesh.geometry.dispose();
          mesh.geometry = new THREE.BoxGeometry(Vector3.x, value, Vector3.z);
        });
      geometryFolder
        .add(Vector3, "z")
        .name("Depth")
        .onChange((value) => {
          mesh.geometry.dispose();
          mesh.geometry = new THREE.BoxGeometry(Vector3.x, Vector3.y, value);
        });
    }

    // Ensure the GUI is opened
    meshFolder.open();
    geometryFolder.open();
  }

  public animate(meshes: THREE.Mesh[]): void {
    const animateLoop = () => {
      requestAnimationFrame(animateLoop);
      if (this.controls) {
        this.controls.update(0.01);
      }
      this.renderer.render(this.scene, this.camera);
    };
    animateLoop();
  }

  public removeAllMeshes(): void {
    this.scene.children.forEach((child) => {
      if (child instanceof THREE.Mesh) {
        this.scene.remove(child);
      }
    });
  }
}

const build3dTilemap = async (mapData: MapBuildData) => {
  const data = mapData;
  const threeDeeScene = new ThreeDeeScene("threedee-view", data.height);
  const tileGeometries: THREE.BufferGeometry[] = [];
  data.tileBoxes.forEach((tile) => {
    // Create a new BoxGeometry
    const newBox = new THREE.BoxGeometry(tile.width, tile.height, tile.depth);

    // Apply the translation (position)
    newBox.translate(tile.position.x, tile.position.y, tile.position.z);

    // Convert the hex color string to a THREE.Color
    const color = new THREE.Color(tile.color.hex); // Use tile's defined color

    // Create a color array for the vertices
    const colors: number[] = [];
    for (let i = 0; i < newBox.attributes.position.count; i++) {
      colors.push(color.r, color.g, color.b); // Push the tile's RGB values for each vertex
    }

    // Set the color attribute for the geometry
    newBox.setAttribute("color", new THREE.Float32BufferAttribute(colors, 3));

    // Push the tile geometry into the array
    tileGeometries.push(newBox);
  });

  // Create a material that supports vertex colors and applies transparency if required
  const material = new THREE.MeshStandardMaterial({
    vertexColors: true,
    wireframe: false,
  });

  // Merge the geometries into one
  const mergedGeometry = BufferGeometryUtils.mergeGeometries(tileGeometries);

  // Add the merged object to the scene
  threeDeeScene.addObject({
    geometry: mergedGeometry,
    material: material,
    position: new THREE.Vector3(0, 0, 0),
    rotation: new THREE.Vector3(0, 0, 0),
  });

  // other meshes just normal boxes
  data.otherMeshes.forEach((tile) => {
    const cube = threeDeeScene.addObject({
      geometry: new THREE.BoxGeometry(tile.width, tile.height, tile.depth),
      material: new THREE.MeshStandardMaterial({
        color: tile.color.hex,
        transparent: tile.color.transparent,
        opacity: tile.color.opacity,
      }),
      position: new THREE.Vector3(
        tile.position.x,
        tile.position.y,
        tile.position.z
      ),
      rotation: new THREE.Vector3(0, 0, 0),
    });
    // threeDeeScene.showMeshGui(cube);
  });

  // Start the animation
  threeDeeScene.animate([]);
};

// Expose the function to the window object
declare global {
  interface Window {
    utils: {
      fetchTileMap: () => void;
    };
  }
}

const build3dLayers = (mapData: MapBuildData) => {
  const data = mapData;

  // create a shape from lines
  const shape = new THREE.Shape();
};

const setup3d = async () => {
  const data = await tileMapData();
  if (!data) {
    throw new Error("Could not fetch tile map data");
  }

  build3dTilemap(data);
  build3dLayers();
};

export { setup3d };
