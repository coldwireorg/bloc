<template>
  <div class="donate-card">
    Help us to keep your service alive !
    <div class="donate-bottom">
      <ButtonAction is-external link="https://coldwire.org/about#donate" text="Donate">
        <Icon name="heart-around" />
      </ButtonAction>
      <a href="#">
        How we spend it
      </a>
    </div>
  </div>
</template>

<style scoped>
  .donate-card {
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 130px;
    width: 190px;
    margin: 0 auto;
    background: var(--primary-green-25);
    border-radius: 8px;
    box-sizing: border-box;
    padding: 1rem;
    font-style: normal;
    font-weight: 600;
    font-size: 14px;
    line-height: 17px;
  }

  .donate-bottom {
    display: flex;
    flex-direction: row;
    gap: 1rem;
    font-size: 12px;
    line-height: 15px;
  }

  .donate-bottom a {
    opacity: 0.5;
  }
</style>
