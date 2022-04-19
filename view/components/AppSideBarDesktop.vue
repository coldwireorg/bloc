<template>
  <div class="main">
    <div class="fx-content">
      <div class="category-container">
        <nuxt-link to="/files" class="category-view">
          <Icon name="folder" />
          My Files
        </nuxt-link>
        <nuxt-link to="/starred" class="category-view">
          <!-- <img src="@/assets/img/icons/star16x16.svg"> -->
          <Icon name="star" />
          Starred
        </nuxt-link>
        <nuxt-link to="/shared" class="category-view">
          <Icon name="share" />
          Shared with me
        </nuxt-link>
      </div>
      <div class="fx-bottom">
        <AccountCardInfo />
        <DonateCard />
      </div>
    </div>
  </div>
</template>

<style scoped>
  .main {
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    width: 254px;
    background: var(--complementary-gray-3);
  }

  .fx-content {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 100vh;
    padding-bottom: 2rem;
    box-sizing: border-box;
    padding-top: 119px;
  }

  .category-container {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    gap: 1rem;
    width: 190px;
    margin: 0 auto;
  }

  .category-view {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    gap: 8px;
    width: 100%;
    height: 32px;
    border-radius: 8px;
    color: white;
    font-family: 'Montserrat';
    font-style: normal;
    font-weight: 600;
    font-size: 12px;
    line-height: 15px;
    padding: 0 8px;
    box-sizing: border-box;
  }

  .category-view:hover {
    background: rgba(240, 246, 252, 0.05);
  }

  .fx-bottom {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  @media screen and (max-height: 580px) {
    .fx-content {
      padding-top: 80px;
    }
  }

  @media screen and (max-height: 500px) {
    .fx-content {
      padding-top: 60px;
    }
  }
</style>
