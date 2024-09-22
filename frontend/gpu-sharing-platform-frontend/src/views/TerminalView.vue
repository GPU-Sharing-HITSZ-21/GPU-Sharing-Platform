<template>
  <div id="app">
    <TerminalComponent />
  </div>
</template>

<script>
import TerminalComponent from '../components/Terminal.vue';

export default {
  name: 'App',
  components: {
    TerminalComponent,
  },
};
</script>

<style>
html, body, #app {
  height: 100%;
  margin: 0;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>