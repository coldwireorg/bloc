<template>
  <div class="fx-column file-card">
    <div class="top-view">
      <div class="fx-row icon-container">
        <Icon name="star-card16x16" type="rounded" padding="5" brd-radius="100" hover />
        <Icon name="option-card16x16" type="rounded" padding="5" brd-radius="100" hover />
      </div>
    </div>
    <div class="fx-column bottom-view">
      <div class="fx-row title-container">
        <Icon name="star" bg-color="#0F7B6C" padding="5" brd-radius="5" />
        <p>{{ fileName }}</p>
      </div>
      <p class="time-history">
        Edited {{ lastEdit }} min. ago
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FileCard',
  props: {
    isFavorite: {
      type: Boolean,
      required: false,
      default: false
    },
    fileName: {
      type: String,
      required: true
    },
    lastEdit: {
      type: String,
      required: true
    },
    mimeType: {
      type: String,
      required: false,
      default: 'application/text'
    }
  },
  data () {
    return {
      time: 5
    }
  },
  methods: {
    getLapsedEditTime () {
      /* */
    }
  }
}
</script>

<style scoped>
  .file-card {
    width: 320px;
    min-width: 320px;
    height: 192px;
    border-radius: 8px;
    background: #30363D;
    box-sizing: border-box;
  }

  .top-view {
    position: relative;
    height: 89px;
    width: 100%;
    padding: 8px;
    background: #0F7B6C;
    border-radius: 8px 8px 0px 0px;
    box-sizing: border-box;
  }

  .icon-container {
    float: right;
  }

  .bottom-view {
    box-sizing: border-box;
    width: 100%;
    padding: 24px 18px;
    gap: 13px;
  }

  .title-container {
    align-items: center;
    gap: 8px
  }

  .time-history {
    font-style: normal;
    font-weight: 600;
    font-size: 12px;
    line-height: 15px;
    color: var(--complementary-white-25);
  }
</style>
