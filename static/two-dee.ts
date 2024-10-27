import { rand } from "three/webgpu";
import { tileMapData } from "./data.ts";
import { MapBuildData, Line } from "./interface.ts";
import { randInt } from "three/src/math/MathUtils.js";

class TwoDeeCanvas {
  private canvas: HTMLCanvasElement;
  private context: CanvasRenderingContext2D;
  private width: number;
  private height: number;
  private dashModifer: number[];

  constructor(canvasId: string, segments: number = 1) {
    const canvas = document.getElementById(canvasId);
    if (!canvas) {
      throw new Error(`Could not find canvas with id ${canvasId}`);
    }
    if (!(canvas instanceof HTMLCanvasElement)) {
      throw new Error(`Element with id ${canvasId} is not a canvas`);
    }
    const container = canvas.parentElement;
    if (!container) {
      throw new Error(
        `Could not find parent element for canvas with id ${canvasId}`
      );
    }
    this.canvas = canvas;

    // Set canvas drawing resolution to match its displayed size
    let minSize = Math.min(container.offsetWidth, container.offsetHeight);

    // Ensure the size is divisible by the number of segments and is an integer
    if (segments > 0) {
      minSize = Math.floor(minSize / segments) * segments;
    }

    this.canvas.width = minSize;
    this.canvas.height = minSize;

    this.setupRandomDashes();

    this.width = this.canvas.width;
    this.height = this.canvas.height;

    const context = canvas.getContext("2d");
    if (!context) {
      throw new Error("Could not get 2d context from canvas");
    }
    this.context = context;
  }

  get canvasWidth() {
    return this.width;
  }

  get canvasHeight() {
    return this.height;
  }

  setupRandomDashes() {
    this.dashModifer = [];
    let dashCount = 0;
    for (let i = 0; i < 10; i++) {
      const dash = randInt(1, 4);
      dashCount += dash;
      this.dashModifer.push(dash);
    }
    console.log("dashModifer", this.dashModifer);
  }

  drawFillRectangle(
    x: number,
    y: number,
    width: number,
    height: number,
    color: string
  ) {
    this.context.fillStyle = color;
    this.context.fillRect(x, y, width, height);
  }

  drawLine(x1: number, y1: number, x2: number, y2: number, color: string) {
    const mod = randInt(1, 2);
    const posmod = randInt(0, 6);
    const dashslice = this.dashModifer.slice(posmod, posmod + mod);
    this.context.strokeStyle = color;
    this.context.beginPath();
    this.context.lineWidth = randInt(1, 2);
    this.context.moveTo(x1, y1);
    this.context.lineTo(x2, y2);
    this.context.setLineDash(dashslice);
    this.context.stroke();
  }

  drawPath(path: Path, color: string) {
    const { lines } = path;
    if (lines.length < 2) {
      return;
    }
    this.context.strokeStyle = color;
    this.context.beginPath();
    this.context.moveTo(lines[0].x, lines[0].y);
    for (let i = 1; i < lines.length; i++) {
      this.context.lineTo(lines[i].x, lines[i].y);
    }
    this.context.stroke();
  }

  clear() {
    this.context.clearRect(0, 0, this.width, this.height);
  }
}

const setup2d = async () => {
  const data = await getMapData();
  build2dTileMap(data);
  stopRenderLoop();
  build2dAnimationLayer(data);
  buildInteractiveGrid();
};

const setup2dLine = async () => {
  const data = await getMapData();
  build2dLineMap(data);
};

let activeRenderId: null | number = null; // Global variable to track the current animation frame ID

const stopRenderLoop = () => {
  if (activeRenderId !== null) {
    cancelAnimationFrame(activeRenderId);
    activeRenderId = null; // Reset to indicate no active render
  }
};

const getMapData = async () => {
  const data = await tileMapData();
  if (!data) {
    throw new Error("Could not fetch tile map data");
  }
  return data;
};

const build2dTileMap = async (
  mapData: MapBuildData,
  loadTime: number = 2000
) => {
  const data = mapData;

  const worldWidth = data.width;
  const worldHeight = data.height;

  const canvas = new TwoDeeCanvas("twodee-view", data.height);
  canvas.clear();

  const tilewidth = canvas.canvasWidth / worldWidth;
  const tileHeight = canvas.canvasHeight / worldHeight;

  canvas.drawFillRectangle(
    0,
    0,
    canvas.canvasWidth,
    canvas.canvasHeight,
    "rgba(255, 255, 255, 0)"
  );

  const midX = Math.floor(worldWidth / 2);
  const midY = Math.floor(worldHeight / 2);

  // Shuffle the tiles to render them in a random order, favoring central tiles
  const shuffledTiles = data.tileBoxes.sort((tilea, tileb) => {
    const distA = Math.sqrt(
      Math.pow(tilea.position.x - midX, 2) +
        Math.pow(tilea.position.z - midY, 2)
    );
    const distB = Math.sqrt(
      Math.pow(tileb.position.x - midX, 2) +
        Math.pow(tileb.position.z - midY, 2)
    );

    // Add a small random factor to keep some randomness in the shuffle
    const randomFactor = Math.random() * 1; // Adjust this factor as needed

    // Sort by distance, but mix in randomness
    return distB - distA + randomFactor;
  });

  const totalTiles = shuffledTiles.length;
  const startTime = performance.now();

  let count = 0;

  // Cubic easing function (ease-in)
  // const easeInCubic = (t) => t * t * t * t;

  // Cubic easing function (ease-out)
  const easeOutCubic = (t) => --t * t * t + 1;

  const renderTile = () => {
    const currentTime = performance.now();
    const elapsedTime = currentTime - startTime;

    // Calculate normalized time (between 0 and 1)
    const t = Math.min(elapsedTime / loadTime, 1); // Ensure t never exceeds 1

    // Apply the cubic easing function to determine the progress
    const easedT = easeOutCubic(t);

    // Calculate how many tiles should be rendered based on the eased time
    const expectedTiles = Math.floor(easedT * totalTiles);

    // Render as many tiles as needed to catch up
    while (count < expectedTiles && shuffledTiles.length > 0) {
      const tile = shuffledTiles.pop();
      if (!tile) {
        return;
      }

      const x = tile.position.x * tilewidth;
      const y = tile.position.z * tileHeight;
      const w = tile.width * tilewidth;
      const h = tile.depth * tileHeight;

      canvas.drawFillRectangle(x, y, w, h, tile.color.hex);
      count++;
    }

    // Stop rendering if all tiles have been rendered
    if (count < totalTiles) {
      requestAnimationFrame(renderTile); // Schedule the next frame
    }
  };

  requestAnimationFrame(renderTile); // Start the rendering loop
};

const buildInteractiveGrid = () => {
  const canvas = new TwoDeeCanvas("twodee-interactive");
  canvas.clear();
};

const hideCanvases = () => {
  const canvases = document.querySelectorAll("canvas");
  canvases.forEach((canvas) => {
    if (canvas instanceof HTMLCanvasElement) {
      canvas.classList.add("-hidden");
    }
  });
};

const revealCanvases = () => {
  const canvases = document.querySelectorAll("canvas");
  canvases.forEach((canvas) => {
    if (canvas instanceof HTMLCanvasElement) {
      canvas.classList.remove("-hidden");
    }
  });
};

const canvasResize = () => {
  const container = document.querySelector<HTMLElement>(".canvas-container");
  if (!container) {
    return;
  }
  const minSize = Math.min(container.offsetWidth, container.offsetHeight);
  const canvases = container.querySelectorAll("canvas");
  canvases.forEach((canvas) => {
    console.log("canvas", canvas);
    if (canvas instanceof HTMLCanvasElement) {
      canvas.width = minSize;
      canvas.height = minSize;
    }
  });
  setup2d();
};

const build2dLineMap = async (mapData: MapBuildData) => {
  const data = mapData;

  const worldWidth = data.width;
  const worldHeight = data.height;

  const canvas = new TwoDeeCanvas("twodeeline-view", data.height);
  canvas.clear();

  const tilewidth = canvas.canvasWidth / worldWidth;
  const tileHeight = canvas.canvasHeight / worldHeight;

  canvas.drawFillRectangle(
    0,
    0,
    canvas.canvasWidth,
    canvas.canvasHeight,
    "rgba(255, 255, 255, 0)"
  );

  const flatLines: Line[] = [];
  Object.values(data.lines).forEach((lineSet) => {
    flatLines.push(...lineSet);
  });
  // shuffle
  flatLines.forEach((line) => {
    const x1 = line.start.x * tilewidth;
    const y1 = line.start.y * tileHeight;
    const x2 = line.end.x * tilewidth;
    const y2 = line.end.y * tileHeight;

    canvas.drawLine(x1, y1, x2, y2, line.color);
  });
};

const build2dAnimationLayer = async (mapData: MapBuildData) => {
  const data = mapData;

  const worldWidth = data.width;
  const worldHeight = data.height;

  const canvas = new TwoDeeCanvas("twodee-animation-layer", data.height);
  canvas.clear();

  const tilewidth = canvas.canvasWidth / worldWidth;
  const tileHeight = canvas.canvasHeight / worldHeight;

  canvas.drawFillRectangle(
    0,
    0,
    canvas.canvasWidth,
    canvas.canvasHeight,
    "rgba(255, 255, 255, 0)"
  );

  const flatLines: Line[] = [];
  Object.values(data.lines).forEach((lineSet) => {
    flatLines.push(...lineSet);
  });
  const shuffledLines = flatLines.sort(() => Math.random() - 0.5);

  const totalLines = shuffledLines.length;
  const startTime = performance.now();

  let count = 0;

  // Cubic easing function (ease-in)
  // const easeInCubic = (t) => t * t * t * t;

  // Cubic easing function (ease-out)
  const easeOutCubic = (t) => --t * t * t + 1;

  const renderLine = () => {
    const currentTime = performance.now();
    const elapsedTime = currentTime - startTime;

    // Calculate normalized time (between 0 and 1)
    const t = Math.min(elapsedTime / 2000, 1); // Ensure t never exceeds 1

    // Apply the cubic easing function to determine the progress
    const easedT = easeOutCubic(t);

    // Calculate how many lines should be rendered based on the eased time
    const expectedLines = Math.floor(easedT * totalLines);

    // Render as many lines as needed to catch up
    while (count < expectedLines && shuffledLines.length > 0) {
      const line = shuffledLines.pop();
      if (!line) {
        return;
      }

      const x1 = line.start.x * tilewidth;
      const y1 = line.start.y * tileHeight;
      const x2 = line.end.x * tilewidth;
      const y2 = line.end.y * tileHeight;

      canvas.drawLine(x1, y1, x2, y2, line.color);
      count++;
    }

    if (count < totalLines) {
      activeRenderId = requestAnimationFrame(renderLine);
    }
  };

  activeRenderId = requestAnimationFrame(renderLine);
};

export {
  TwoDeeCanvas,
  build2dTileMap,
  buildInteractiveGrid,
  setup2d,
  setup2dLine,
  canvasResize,
  hideCanvases,
  revealCanvases,
};
