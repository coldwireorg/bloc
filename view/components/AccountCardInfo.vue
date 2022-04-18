<template>
  <nuxt-link to="/settings" class="fx-column account-card-info">
    <div class="fx-row infos">
      <div class="fx-row userinfo">
        <div class="profile-icon">
          <h6>J</h6>
        </div>
        <h4>John Doe</h4>
      </div>
      <nuxt-link to="/">
        <Icon name="disconnect" padding="5" hover brd-radius="100" />
        <!-- <img src="@/assets/img/icons/disconnect16x16.svg" alt="Disconnect"> -->
      </nuxt-link>
    </div>
    <div class="fx-row storage-quota">
      <ProgressBar type="determinate" fill-level="33" />
      350Mo/1Go
    </div>
  </nuxt-link>
</template>

<style scoped>
  .account-card-info {
    position: relative;
    justify-content: space-between;
    height: 90px;
    width: 190px;
    margin: 0 auto;
    background: var(--complementary-white-5);
    border-radius: 8px;
    padding: 1rem;
    box-sizing: border-box;
    transition: 0.15s ease-in-out;
  }

  .account-card-info:hover {
    box-shadow: 0px 0px 2.4px 0px rgba(255, 255, 255, 0.623);
  }

  .infos {
    justify-content: space-between;
    align-items: center;
  }

  .userinfo {
    justify-content: flex-start;
    align-items: center;
    gap: 1rem;
  }

  .profile-icon {
    position: relative;
    width: 32px;
    height: 32px;
    background: #0F7B6C;
    border-radius: 100%;
    text-align: center;
  }

  .profile-icon h6 {
    position: absolute;
    font-size: 12px;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
  }

  .disconnect-button {
    width: 16px;
    height: 16px;
    padding: 5px;
    border-radius: 100%;
    transition: 0.1s ease-in-out;
  }

  .disconnect-button:hover {
    background: rgba(255, 255, 255, 0.055);
  }

  .storage-quota {
    align-items: center;
    justify-content: center;
    gap: 8px;
    font-size: 8px;
  }

  h4, h5, h6 {
    margin: 0;
  }
</style>
