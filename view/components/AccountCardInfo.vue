<template>
  <div class="account-card-info">
    <div class="infos">
      <div class="userinfo">
        <div class="profile-icon">
          <h6>J</h6>
        </div>
        <h4>John Doe</h4>
      </div>
      <img src="@/assets/img/disconnect16x16.svg" alt="Disconnect">
    </div>
    <div class="storage-quota">
      <ProgressBar />
      350Mo/1Go
    </div>
  </div>
</template>

<style scoped>
  .account-card-info {
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 90px;
    width: 190px;
    margin: 0 auto;
    background: var(--complementary-white-5);
    border-radius: 8px;
    padding: 1rem;
    box-sizing: border-box;
  }

  .infos {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }

  .userinfo {
    display: flex;
    flex-direction: row;
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

  .storage-quota {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 8px;
    font-size: 8px;
  }

  h4, h5, h6 {
    margin: 0;
  }
</style>
