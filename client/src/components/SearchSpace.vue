<template>
  <div id="searchSpace">
    <h1>検索するよ〜</h1>
    <input v-model="inputValue" placeholder="検索内容を入力しろ"/>
    <p>{{inputValue}}</p>
    <!-- エラーになるためボタンの記述はコメントアウト -->
    <!-- <div @click="search"><v-btn class="button">検索</v-btn></div> -->
  </div>
</template>

<script>
  export default {
    // クリック時のsearchの挙動を記述してないため後で書く
    name: 'SearchSpace',
    data () {
      return {
        inputValue: ""
      }
    }
  }
</script>

<style scoped>
/* デザインは後からvuetifyとかで良い感じにする */
  input {
    padding: 10px;
  }
  .button {
    font-size: 20px;
    border: 1px solid ;
    background-color: salmon;
    font-family: 'Courier New', Courier, monospace;
    border-radius: 10px;
    padding: 3px 15px;
  }
</style>
