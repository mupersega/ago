import * as THREE from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";

import GUI from 'lil-gui'

interface ThreeDeeObject {
    geometry: THREE.BufferGeometry;
    material: THREE.Material;
    position?: THREE.Vector3;
    rotation?: THREE.Vector3;
}

interface TileBox {
    depth: number;
    height: number;
    width: number;
    position: {x: number, y: number, z: number}
    color: {
        hex: string;
        transparent: boolean;
        opacity: number;
    };
}

class ThreeDeeScene {
    private scene: THREE.Scene;
    private camera: THREE.PerspectiveCamera;
    private renderer: THREE.WebGLRenderer;
    private displayElement: HTMLElement;
    private controls: OrbitControls;
    private directionalLight: THREE.DirectionalLight;
    private gui: GUI; // Correctly typed GUI instance

    constructor(private displayId: string) {
        this.scene = new THREE.Scene();
        const display = document.getElementById(this.displayId);

        if (!display) {
            throw new Error(`No element with id "${this.displayId}" found`);
        }
        this.displayElement = display;

        this.setupRenderer();
        this.setupCamera();
        this.addLights();
        this.setupControls();
        this.addHelpers();
        this.setupGUI();
        this.handleWindowResize();
    }

    private setupRenderer(): void {
        const { clientWidth, clientHeight } = this.displayElement;
        this.renderer = new THREE.WebGLRenderer({ canvas: this.displayElement });
        this.renderer.setSize(clientWidth, clientHeight);
    }

    private setupCamera(): void {
        const { clientWidth, clientHeight } = this.displayElement;
        this.camera = new THREE.PerspectiveCamera(
            75,
            clientWidth / clientHeight,
            0.1,
            1000
        );
        this.camera.position.set(25, 25, 50);
        this.camera.lookAt(0, 0, 0);
    }

    private setupControls(): void {
        this.controls = new OrbitControls(this.camera, this.renderer.domElement);
        this.controls.enableDamping = true;
        this.controls.dampingFactor = 0.05;
        this.controls.screenSpacePanning = false;
        this.controls.maxPolarAngle = Math.PI / 2;
    }

    private addLights(): void {
        this.directionalLight = new THREE.DirectionalLight(0xffffff, 1);
        this.directionalLight.position.set(10, 10, 5); 
        this.directionalLight.castShadow = true; 

        this.scene.add(this.directionalLight);

        const directionalLight2 = new THREE.DirectionalLight(0xffffff, 1);
        directionalLight2.position.set(-10, 10, -5);
        directionalLight2.castShadow = true;
        this.scene.add(directionalLight2);
    }

    private addHelpers(): void {
        const axesHelper = new THREE.AxesHelper(5);
        this.scene.add(axesHelper);

        const size = 30;
        const divisions = 30;
        const gridHelper = new THREE.GridHelper(size, divisions);
        this.scene.add(gridHelper);
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
        this.gui.add( document, 'title' );

        // // Add basic parameters to the GUI
        // const cameraFolder = this.gui
        // cameraFolder.add(this.camera.position, "x", -50, 50).name("X Position");
        // cameraFolder.add(this.camera.position, "y", -50, 50).name("Y Position");
        // cameraFolder.add(this.camera.position, "z", -50, 50).name("Z Position");
        // cameraFolder.open();
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
            geometryFolder.add(Vector3, "x").name("Width").onChange((value) => {
                mesh.geometry.dispose();
                mesh.geometry = new THREE.BoxGeometry(value, Vector3.y, Vector3.z);
            }
            );
            geometryFolder.add(Vector3, "y").name("Height").onChange((value) => {
                mesh.geometry.dispose();
                mesh.geometry = new THREE.BoxGeometry(Vector3.x, value, Vector3.z);
            }
            );
            geometryFolder.add(Vector3, "z").name("Depth").onChange((value) => {
                mesh.geometry.dispose();
                mesh.geometry = new THREE.BoxGeometry(Vector3.x, Vector3.y, value);
            }
            );
        }

        // Ensure the GUI is opened
        meshFolder.open();
        geometryFolder.open();
    }

    public animate(meshes: THREE.Mesh[]): void {
        const animateLoop = () => {
            requestAnimationFrame(animateLoop);
            this.controls.update(); 
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
const addWaterCube = (scene: ThreeDeeScene) => {
    const cube = scene.addObject({
        geometry: new THREE.BoxGeometry(100, 8.5, 100),
        material: new THREE.MeshStandardMaterial({ color: 0x0000bb, transparent: true, opacity: 0.5 }),
        position: new THREE.Vector3(0, 0, 0),
        rotation: new THREE.Vector3(0, 0, 0),
    });
    scene.showMeshGui(cube);
}

const FetchTileMap = () => {
    fetch("/tilemap")
        .then((response) => response.json())
        .then((data: TileBox[]) => {
            const threeDeeScene = new ThreeDeeScene("three-dee");
            console.log('okay')
            let count = 0;
            data.forEach((tile) => {
                const newCube = threeDeeScene.addObject({
                    geometry: new THREE.BoxGeometry(tile.width, tile.height, tile.depth),
                    material: new THREE.MeshStandardMaterial({color: tile.color.hex, transparent: tile.color.transparent, opacity: tile.color.opacity}),
                    position: new THREE.Vector3(tile.position.x, tile.position.y, tile.position.z),
                    rotation: new THREE.Vector3(0, 0, 0),
                });
                count += 1;
                if (count === data.length) {
                    threeDeeScene.showMeshGui(newCube);
                }
            });
            threeDeeScene.animate([]);
        });
}

const toggleThreeDeeMode = () => {
    const threeDeeElement = document.getElementById("three-dee");
    if (threeDeeElement) {
        threeDeeElement.classList.toggle("hidden");
    }
}

// Expose the function to the window object
declare global {
    interface Window {
        utils: {
            fetchTileMap: () => void;
        };
    }
}

window.utils = {
    fetchTileMap: FetchTileMap,
};