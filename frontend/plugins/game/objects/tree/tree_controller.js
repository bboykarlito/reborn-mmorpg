import TreeObserver from "~/plugins/game/objects/tree/tree_observer";
import TreeState from "~/plugins/game/objects/tree/tree_state";
import { EventBus } from "~/plugins/game/event_bus";

class TreeController {
  constructor(gameObject) {
    this.state = new TreeState(gameObject);
    this.observer = new TreeObserver(this.state);
  }

  remove() {
    this.state = null
    this.observer.remove()
  }
}

export default TreeController;
