<template>
  <div>
    <a v-if="isExternal" target="_blank" :href="link" class="action-button">
      {{ text }}
      <slot />
    </a>
    <nuxt-link v-else :to="link" class="action-button">
      {{ text }}
      <slot />
    </nuxt-link>
  </div>
</template>

<script>
export default {
  name: 'ButtonAction',
  props: {
    text: {
      type: String,
      required: true
    },
    link: {
      type: String,
      required: false,
      default: '#'
    },
    isExternal: {
      type: Boolean,
      required: false,
      default: false
    }
  }
}
</script>

<style scoped>
  .action-button {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    box-sizing: border-box;
    padding: 8px 8.5px;
    gap: 6px;
    background: var(--primary-green);
    border-radius: 8px;
  }
</style>
