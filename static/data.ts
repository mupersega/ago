import { MapBuildData } from "./interface";

export const tileMapData = async (): Promise<MapBuildData> => {
  const response = await fetch(`/tilemap?${new Date().getTime()}`);
  console.log(response);
  return response.json();
};

declare global {
  interface Window {
    getTileMap: () => Promise<MapBuildData>;
  }
}

window.getTileMap = tileMapData;
