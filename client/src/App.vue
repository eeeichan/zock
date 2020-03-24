<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <SearchZozo :items="items"></SearchZozo> <!-- :はv-bindの省略らしい -->
  </div>
</template>

<script>
import SearchZozo from './components/SearchZozo.vue'

export default {
  name: 'App',
  components: {
    SearchZozo
  },
  // data () { でも問題なかった
  data: () => {
    return {
      // 仮データ。イメージ。
      items: [
        {
          id: 1,
          genre: 'accessories',
          brand: 'Maison Margiela',
          name: 'Maison Margiela スモールウォレット',
          color: 'グレー',
          size: '',
          price: '¥53,900',
          salePrice: '¥48,510',
          saleRate1: '10%',
          saleRate2: 'なし',
          lowDate: '2020/3/17',
          stock: true,
        },
        {
          id: 2,
          genre: 'accessories',
          brand: 'Maison Margiela',
          name: 'Maison Margiela スモールウォレット',
          color: 'ブラック',
          size: '',
          price: '¥53,900',
          salePrice: '¥48,510',
          saleRate1: '10%',
          saleRate2: 'なし',
          lowDate: '2020/3/17',
          stock: false,
        },
        {
          id: 3,
          genre: 'accessories',
          brand: 'Maison Margiela',
          name: 'Maison Margiela スモールウォレット',
          color: 'ライトピンク',
          size: '',
          price: '¥53,900',
          salePrice: '¥48,510',
          saleRate1: '10%',
          saleRate2: 'なし',
          lowDate: '2020/3/17',
          stock: false,
        },
        {
          id: 4,
          genre: 'shoes',
          brand: 'Maison Margiela',
          name: 'Maison Margiela タビバレエシューズ',
          color: 'ブラック',
          size: '35',
          price: '¥96,800',
          salePrice: '¥95,800',
          saleRate1: 'なし',
          saleRate2: '1000円割引',
          lowDate: '2020/3/17',
          stock: false,
        },
        {
          id: 5,
          genre: 'shoes',
          brand: 'Maison Margiela',
          name: 'Maison Margiela タビバレエシューズ',
          color: 'ベージュ',
          size: '35',
          price: '¥96,800',
          salePrice: '¥95,800',
          saleRate1: 'なし',
          saleRate2: '1000円割引',
          lowDate: '2020/3/17',
          stock: false,
        }
      ]
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
