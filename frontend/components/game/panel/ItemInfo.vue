<template>
  <div id="item_info-panel" class="game-panel" v-if="showItemInfoPanel">
    <div class="game-panel-content">
      <h4 class="heading">{{ itemInfo["kind"] }}</h4>
      <div v-if="itemInfo['crafted_by']">
        Crafted by: {{ itemInfo["crafted_by"]["name"] }}
      </div>
      <div v-if="itemInfo['payed_until']">
        Payed until: {{ new Date(itemInfo["payed_until"]) }}
      </div>
      <button type="button" class="rpgui-button" @click="showItemInfoPanel = false"><p>Close</p></button>
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      showItemInfoPanel: false,
      itemInfo: {},
    }
  },

  created() {
    EventBus.$on("item_info", this.showItemInfo)
  },

  beforeDestroy() {
    EventBus.$off("item_info", this.showItemInfo)
  },

  methods: {
    showItemInfo(data) {
      this.showItemInfoPanel = true
      this.itemInfo = data
    },
  }
}
</script>


<style lang="scss">
#item_info-panel {
  position: absolute;
  top: 300px;
  left: 350px;
  color: white;
  .heading {
    margin-top: 0px;
  }
}
.game-panel-content {
  div {
    color: white;
    font-size: 11px;
  }
}
</style>
