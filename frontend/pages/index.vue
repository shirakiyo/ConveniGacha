<template>
  <v-layout>
    <div class="d-flex justify-center flex-column ma-auto">
      <div class="d-flex justify-space-around ma-auto" style="width: 80vw; max-width: 600px;">
        <div><input type="radio" v-model="category" value="" class="mr-2" checked><span style="color: #3995D8;font-weight: bold;">全商品</span></div>
        <div><input type="radio" v-model="category" value="foods" class="mr-2"><span style="color: #3995D8;font-weight: bold;">食糧</span></div>
        <div><input type="radio" v-model="category" value="sweets" class="mr-2"><span style="color: #3995D8;font-weight: bold;">デザート</span></div>
        <div><input type="radio" v-model="category" value="snacks" class="mr-2"><span style="color: #3995D8;font-weight: bold;">お菓子</span></div>
      </div>
      <v-btn @click="getProduct" color="#2EAB4F" class="mt-6 mb-8 ma-auto gacha__btn" style="color: white;">ガチャを回す</v-btn>
      <v-card v-show="product.name" max-width="800px" style="width: 90vw;">
        <v-card-text class="pa-8 d-flex justify-around flex-column">
          <h2 align="center" style="font-size: 1.2rem; color: black;">{{product.name}}</h2>
          <span class="mt-2" style="max-width: 600px; width: 100%; margin: 0 auto; height: 5px; background: #2EAB4F"></span>
          <p class="mt-4">{{product.detail}}</p>
          <div class="d-flex flex-column justify-center">
            <h3 align="center">{{product.price}}円(税込)</h3>
            <a :href="product.link" target="_blank" ref="noopener" class="d-flex align-center justify-center mt-4" color="blue" style="text-decoration: none; color: #2196F3;">
              ファミマ公式サイトでチェック
              <v-icon color="blue">mdi-open-in-new</v-icon>
            </a>
          </div>
        </v-card-text>
      </v-card>
      <div v-show="productHistory.length != 0" align="center" class="mt-8">
        <p @click.stop="dialog = true" style="font-weight: bold; border-bottom: 2px solid; display: inline; cursor: pointer;">履歴を確認する</p>
      </div>

      <v-dialog v-model="dialog" max-width="600">
        <v-card>
          <v-card-title style="background: #3995D8; color: white;">
            ガチャ履歴
            <v-spacer></v-spacer>
            <v-icon color="white" @click="dialog=false">mdi-close</v-icon>
          </v-card-title>
          <v-card-text class="product__history-items">
            <v-list v-for="(p, i) in productHistory" class="product__history-item" :key="p.name">
              <a :href="p.link" style="color: gray;" target="_blank" rel="noopener">{{p.name}}</a>
              <span align="center">{{p.price}}円</span>
              <v-icon color="#ff4b3e" @click="deleteProductFromHistory(i)" style="cursor: pointer;">mdi-delete</v-icon>
            </v-list>
          </v-card-text>
          <v-card-actions class="py-4">
            <v-btn color="red darken-1" @click="resetHistory" text>リセット</v-btn>
            <v-spacer></v-spacer>
            <span>合計金額： {{totalPrice}}円</span>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-layout>
</template>

<script>
  import axios from "axios"

export default {
    data() {
      return {
        product: {},
        category: "",
        productHistory: [],

        dialog: false,
        totalPrice: 0
      }
    },
    methods: {
      getProduct() {
        axios.get(
          process.env.NUXT_ENV_API_ENDPOINT + "/familyMart/get",
          {
            params: {
              category: this.category
            }
          }
        )
        .then(res => {
          this.product = res.data
          this.productHistory.unshift(res.data)
          this.totalPrice += res.data.price
        })
        .catch(err => {
          console.log(err)
        })
      },
      resetHistory() {
        this.totalPrice = 0
        this.productHistory = []
        this.dialog = false
      },
      deleteProductFromHistory(i) {
        this.totalPrice -= this.productHistory[i].price
        this.productHistory.splice(i, 1)
      }
    },
}
</script>

<style scoped>
  .gacha__btn {
    max-width: 400px;
    font-size: 1.2rem !important;
    font-weight: bold;
    display: flex;
    align-items: center;
  }

  .product__history-items {
    border-bottom: 2px solid #3995D8;
    max-height: 60vh;
    overflow: scroll;
    margin-top: 20px;
  }

  .product__history-item {
    display: grid;
    align-items: center;
    grid-template-columns: 10fr 3fr 1fr;
    border-bottom: 1px dashed #3995D8 !important;
  }

  .product__history-item:last-child {
    border-bottom: none !important;
  }
</style>
