import { tileMapData } from "./data.ts";

class TwoDeeCanvas {
  private canvas: HTMLCanvasElement;
  private context: CanvasRenderingContext2D;
  private width: number;
  private height: number;

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

  clear() {
    this.context.clearRect(0, 0, this.width, this.height);
  }
}

const setup2d = () => {
  build2dTileMap();
  buildInteractiveGrid();
};

let activeRenderId: null | number = null; // Global variable to track the current animation frame ID

const build2dTileMap = async (loadTime: number = 1000) => {
  const data = await tileMapData();
  if (!data) {
    throw new Error("Could not fetch tile map data");
  }

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
    "#fff"
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

export {
  TwoDeeCanvas,
  build2dTileMap,
  buildInteractiveGrid,
  setup2d,
  canvasResize,
  hideCanvases,
  revealCanvases,
};
