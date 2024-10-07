import { build2dTileMap, setup2d } from "./two-dee";
import { setup3d } from "./three-dee";

type Listener<T> = (newState: T, oldState: T) => void;

interface AppState {
  user: string;
  viewMode: string;
}

class StateManager<T extends object> {
  private state: T;
  private listeners: Listener<T>[] = [];

  constructor(initialState: T) {
    this.state = initialState;
  }

  // Method to get the current state
  getState(): T {
    return this.state;
  }

  // Method to set the entire state
  setState(newState: T): void {
    const oldState = { ...this.state }; // Capture the old state
    this.state = newState;
    this.notifyListeners(oldState);
  }

  // Method to update a subset of the state
  update(partialState: Partial<T>): void {
    const oldState = { ...this.state }; // Capture the old state
    this.state = { ...this.state, ...partialState };
    this.notifyListeners(oldState);
  }

  // Method to subscribe to state changes
  subscribe(listener: Listener<T>): void {
    this.listeners.push(listener);
  }

  // Method to unsubscribe from state changes
  unsubscribe(listener: Listener<T>): void {
    this.listeners = this.listeners.filter((l) => l !== listener);
  }

  // Notify all listeners about the state change
  private notifyListeners(oldState: T): void {
    this.listeners.forEach((listener) => listener(this.state, oldState));
  }
}

interface AppState {
  user: string;
  viewMode: string;
}

const appState = new StateManager<AppState>({
  user: "Guest",
  viewMode: "2d",
});

// Update the state
appState.setState({ user: "Admin", viewMode: "2d" });

// Subscribe to changes
appState.subscribe((state: AppState, oldState: AppState) => {
  if (state.user !== oldState.user) {
    console.log(`User changed from ${oldState.user} to ${state.user}`);
  }
  processViewModeChange(state, oldState);
  if (state.viewMode !== oldState.viewMode) {
  }
  updateGui(state, oldState);
});

function processViewModeChange(newState: AppState, oldState: AppState) {
  console.log(
    `View mode changed from ${oldState.viewMode} to ${newState.viewMode}`
  );
  const event = new CustomEvent("viewModeChanged", {
    detail: { state: newState },
  });

  window.dispatchEvent(event);

  if (newState.viewMode === "2d") {
    setup2d();
  } else {
    setup3d();
  }
}

function updateGui(newState: AppState, oldState: AppState) {
  const userElement = document.getElementById("user");
  if (userElement) {
    userElement.value = newState.user;
  }

  const modeElement = document.getElementById("mode");
  if (modeElement) {
    modeElement.value = newState.viewMode;
  }
}

export { appState };
