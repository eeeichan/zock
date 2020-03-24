<template>
  <div id="search">
    <h1>検索結果</h1>
    <table div="table">
      <tr>
        <th>ID</th>
        <th>GENRE</th>
        <th>BRAND</th>
        <th>NAME</th>
        <th>COLOR</th>
        <th>SIZE</th>
        <th>PRICE</th>
        <th>SALEPRICE</th>
        <th>SALERATE1</th>
        <th>SALERATE2</th>
        <th>LOWDATE</th>
        <th>STOCK</th>
      </tr>
      <tr v-for="(item, key) in items" :key="key">
        <td v-text="item.id"></td>
        <td v-text="item.genre"></td>
        <td v-text="item.brand"></td>
        <td v-text="item.name"></td>
        <td v-text="item.color"></td>
        <td v-text="item.size"></td>
        <td v-text="item.price"></td>
        <td v-text="item.salePrice"></td>
        <td v-text="item.saleRate1"></td>
        <td v-text="item.saleRate2"></td>
        <td v-text="item.lowDate"></td>
        <td v-text="item.stock"></td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: 'SearchZozo',
  props: {
    items: String
  }
}
</script>

<style scoped>
  table {
    border: solid;
    margin: 0 auto;
  }
  th,td {
    text-align: center;
    padding: 10px;
  }
  table, tr, th, td {
    border: 0.5px solid;
    border-collapse: collapse;
  }
</style>