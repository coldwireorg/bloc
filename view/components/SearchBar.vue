<template>
  <div class="search-bar">
    <input type="text" placeholder="Search">
  </div>
</template>

<style scoped>
  .search-bar {
    width: 192px;
    background: var(--complementary-gray-1);
    border: 1px solid var(--complementary-white-25);
    height: 2rem;
    box-sizing: border-box;
    border-radius: 8px;
  }

  .search-bar input {
    width: 100%;
    height: 100%;
    background: transparent;
    outline: none;
    border: none;
    padding: 0 13px;
    color: rgb(117, 117, 117);
  }
</style>
