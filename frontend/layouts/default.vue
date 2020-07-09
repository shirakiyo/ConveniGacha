<template>
  <v-app style="background: #f3f4f3;">
    <v-content>
      <header class="conveni__head">
        <div></div>
        <h1 class="d-flex align-center justify-center flex-column" style="font-size: 2.2rem; text-align: center; padding: 1rem 0;">
          <div><span style="color: #2EAB4F;">FamilyMart</span><span style="font-size: 2rem;color: #3995D8;">ガチャ</span></div>
        </h1>
        <div class="d-flex">
          <v-btn class="conveni__btn button" style="border: 2px solid #F56500; color: #F56500; font-weight: bold;"><img class="mr-2" src="../static/seven_icon.png" width="30">ガチャ</v-btn>
          <v-btn class="conveni__btn button" style="border: 2px solid #0687d1; color: #0687d1; font-weight: bold;"><img class="mr-2" src="../static/lawson_icon.jpg" widht="30" height="30">ガチャ</v-btn>
        </div>
      </header>
      <v-container style="min-height: calc(100vh - 184px);">
        <nuxt />
      </v-container>
      <footer class="conveni__foot">
        <div class="d-flex justify-space-around py-6 ma-auto" style="max-width: 600px;">
          <v-btn class="button" style="border: 2px solid #F56500; color: #F56500; font-weight: bold;"><img class="mr-2" src="../static/seven_icon.png" width="30">ガチャ</v-btn>
          <v-btn class="button" style="border: 2px solid #0687d1; color: #0687d1; font-weight: bold;"><img class="mr-2" src="../static/lawson_icon.jpg" widht="30" height="30">ガチャ</v-btn>
        </div>
      </footer>
    </v-content>
  </v-app>
</template>

<script>
export default {
}
</script>

<style scoped>
  .v-application {
    font-family: 'Noto Sans JP', sans-serif !important;
  }

  body {
    background: #f3f4f3 !important;
  }

  .conveni__foot {
    background: white;
    height: 100px;
  }

  .conveni__btn {
    display: none;
    margin: 0 1rem;
  }

  .button {
    padding: 0 24px !important;
    height: 52px !important;
  }

  .conveni__head {
    box-shadow: 0px 3px 10px 5px #f3f4f3;
    background: white;
    display: grid;
    grid-template-columns: 0 3fr 0;
    align-items: center;
  }

  @media screen and (min-width: 600px) {
    .conveni__btn {
      display: flex;
    }

    .conveni__foot {
      display: none;
    }

    .conveni__head {
      grid-template-columns: 1fr 3fr 1fr;
    }
  }
</style>
