interface MapBuildData {
  tileBoxes: TileBox[];
  otherMeshes: TileBox[];
  width: number;
  height: number;
  lines: { [key: number]: Line[] };
}

interface TileBox {
  depth: number;
  height: number;
  width: number;
  position: { x: number; y: number; z: number };
  color: {
    hex: string;
    transparent: boolean;
    opacity: number;
  };
}

interface Vector2 {
  x: number;
  y: number;
}

interface Path {
  lines: Vector2[];
}

interface Line {
  start: Vector2;
  end: Vector2;
  color: string;
}

export { MapBuildData, TileBox, Vector2, Path, Line };
