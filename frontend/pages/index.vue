<template>
  <v-layout>
    <div class="d-flex justify-center flex-column ma-auto">
      <v-btn @click="getProduct" color="white" class="pa-8 mb-16 ma-auto gacha__btn">ガチャを回す</v-btn>
      <v-card v-show="products.name" max-width="800px" style="width: 90vw;">
        <v-card-text class="pa-8 d-flex justify-around flex-column">
          <h2 align="center" style="color: black;">{{products.name}}</h2>
          <div class="ma-auto my-4" style="width:60%;">
            <img :src="products.imageURL" style="width: 100%;">
          </div>
          <div class="d-flex flex-column justify-center">
            <h3 align="center">{{products.price}}円(税込)</h3>
            <a :href="products.link" target="_blank" ref="noopener" class="d-flex align-center justify-center mt-4" color="blue" style="text-decoration: none; color: #2196F3;">
              サイトでチェック
              <v-icon color="blue">mdi-open-in-new</v-icon>
            </a>
          </div>
        </v-card-text>
      </v-card>
    </div>
  </v-layout>
</template>

<script>
  import axios from "axios"

export default {
    data() {
      return {
        products: {}
      }
    },
    methods: {
      getProduct() {
        this.products = ""
        axios.get(
          process.env.NUXT_ENV_API_ENDPOINT + "/get"
        )
        .then(res => {
          this.products = res.data
        })
        .catch(err => {
          console.log(err)
        })
      }
    },
}
</script>

<style>
  .gacha__btn {
    max-width: 400px;
    font-size: 1.2rem !important;
    font-weight: bold;
  }
</style>
