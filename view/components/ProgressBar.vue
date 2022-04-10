<template>
  <div class="progress-bar">
    <div v-if="type == 'determinate'" class="determinate" :class="getClassTheme()" :style="{ width: fillLevel + '%' }" />
    <div v-else class="indeterminate" :class="getClassTheme()" />
  </div>
</template>

<script>
export default {
  name: 'ProgressBar',
  props: {
    type: {
      type: String,
      required: true,
      default: 'indeterminate'
    },
    fillLevel: {
      type: String,
      required: true
    },
    theme: {
      type: String,
      required: false,
      default: 'light'
    }
  },
  methods: {
    getClassTheme () {
      // eslint-disable-next-line no-console
      console.log(this.theme)
      if (this.theme === 'green') {
        return 'thm-green'
      } else {
        return 'thm-light'
      }
    }
  }
}
</script>

<style scoped>
  .progress-bar {
    position: relative;
    height: 4px;
    display: block;
    width: 100%;
    background-color: var(--complementary-white-25);
    /* background-color: #acece6; */
    border-radius: 2px;
    background-clip: padding-box;
    margin: 0.5rem 0 1rem 0;
    overflow: hidden;
    margin: auto 0;
  }

  .progress-bar .determinate {
    position: absolute;
    top: 0;
    bottom: 0;
    transition: width .3s linear;
  }

  .progress-bar .indeterminate:before {
    content: '';
    position: absolute;
    background-color: inherit;
    top: 0;
    left: 0;
    bottom: 0;
    will-change: left, right;
    -webkit-animation: indeterminate 2.1s cubic-bezier(0.65, 0.815, 0.735, 0.395) infinite;
    animation: indeterminate 2.1s cubic-bezier(0.65, 0.815, 0.735, 0.395) infinite;
    }

  .progress-bar .indeterminate:after {
    content: '';
    position: absolute;
    background-color: inherit;
    top: 0;
    left: 0;
    bottom: 0;
    will-change: left, right;
    -webkit-animation: indeterminate-short 2.1s cubic-bezier(0.165, 0.84, 0.44, 1) infinite;
            animation: indeterminate-short 2.1s cubic-bezier(0.165, 0.84, 0.44, 1) infinite;
    -webkit-animation-delay: 1.15s;
            animation-delay: 1.15s;
  }

  .thm-light {
    background-color: #F0F6FC;
  }

  .thm-green {
    background-color: #26a69a;
  }

  @-webkit-keyframes indeterminate {
    0% {
      left: -35%;
      right: 100%; }
    60% {
      left: 100%;
      right: -90%; }
    100% {
      left: 100%;
      right: -90%; }
  }
  @keyframes indeterminate {
    0% {
      left: -35%;
      right: 100%; }
    60% {
      left: 100%;
      right: -90%; }
    100% {
      left: 100%;
      right: -90%; }
  }
  @-webkit-keyframes indeterminate-short {
    0% {
      left: -200%;
      right: 100%; }
    60% {
      left: 107%;
      right: -8%; }
    100% {
      left: 107%;
      right: -8%; }
  }
  @keyframes indeterminate-short {
    0% {
      left: -200%;
      right: 100%; }
    60% {
      left: 107%;
      right: -8%; }
    100% {
      left: 107%;
      right: -8%; }
  }
</style>
