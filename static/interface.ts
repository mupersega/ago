interface MapBuildData {
  tileBoxes: TileBox[];
  otherMeshes: TileBox[];
  width: number;
  height: number;
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

export { MapBuildData, TileBox };
