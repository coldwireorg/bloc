<template>
  <Icon
    v-if="type === 'binder'"
    name="binder-white32x32"
    bg-color="var(--secondary-green)"
    size="10"
    padding="5"
    margin="2"
    brd-radius="5"
  />
  <Icon
    v-else-if="type === 'document'"
    name="document-white32x32"
    bg-color="var(--secondary-blue)"
    size="10"
    padding="5"
    margin="2"
    brd-radius="5"
  />
  <Icon
    v-else-if="type === 'presentation'"
    name="file-white32x32"
    bg-color="var(--secondary-red)"
    size="10"
    padding="5"
    margin="2"
    brd-radius="5"
  />
  <Icon
    v-else-if="type === 'image'"
    name="image-white32x32"
    bg-color="var(--secondary-brown)"
    size="10"
    padding="5"
    margin="2"
    brd-radius="5"
  />
  <Icon
    v-else-if="type === 'video'"
    name="video-white32x32"
    bg-color="var(--secondary-orange)"
    size="10"
    padding="5"
    margin="2"
    brd-radius="5"
  />
  <Icon
    v-else-if="type === 'audio'"
    name="music-white32x32"
    bg-color="var(--secondary-violet)"
    size="10"
    padding="5"
    margin="2"
    brd-radius="5"
  />
  <Icon
    v-else-if="type === 'folder'"
    name="folder-explorer24x24"
    size="24"
  />
  <Icon
    v-else
    name="file-white32x32"
    bg-color="var(--secondary-red)"
    size="10"
    padding="5"
    margin="2"
    brd-radius="5"
  />
</template>

<script>
export default {
  name: 'IconFile',
  props: {
    type: {
      type: String,
      required: true
    }
  }
}
</script>
