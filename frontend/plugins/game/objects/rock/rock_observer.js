import { EventBus } from "~/plugins/game/event_bus";
import Atlas from "~/plugins/game/atlas/atlas";
import GameObserver from "~/plugins/game/game_observer";

class RockObserver {
  constructor(state) {
    this.scene = null;
    this.state = state;
    this.container = null;
    this.mesh = null;
    this.meshRotation = 0
    if (GameObserver.loaded) {
      this.scene = GameObserver.scene;
      this.create();
    } else {
      EventBus.$on("scene-created", scene => {
        this.scene = scene;
        this.create();
      });
    }

  }

  create() {
    let mesh = Atlas.get(this.state.kind + "Rock").clone("rock-" + this.state.id)
    mesh.setParent(null)
    mesh.setEnabled(true);
    mesh.name = "rock-" + this.state.id;
    mesh.position.x = this.state.x
    mesh.position.y = 0
    mesh.position.z = this.state.y
    if (this.state.rotation) {
      let rotationDelta = this.meshRotation - this.state.rotation;
      if (rotationDelta != 0) {
        this.meshRotation = this.state.rotation;
        mesh.rotate(BABYLON.Axis.Y, rotationDelta);
      }
    }
    mesh.metadata = {
      x: this.state.x,
      y: this.state.y,
      id: this.state.id,
      state: this.state
    };
    mesh.setEnabled(true);
    mesh.freezeWorldMatrix();
    mesh.doNotSyncBoundingInfo = true;
    this.mesh = mesh;
  }

  remove() {
    EventBus.$off("scene-created", this.sceneCreatedCallback);
    this.mesh.dispose();
    this.mesh = null;
    this.state = null;
  }
}

export default RockObserver;
