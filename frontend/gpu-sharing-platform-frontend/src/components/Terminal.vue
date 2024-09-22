<template>
  <div id="terminal-container">
    <!-- Terminal will be rendered in this div -->
    <div id="terminal"></div>
  </div>
</template>

<script>
import { Terminal } from 'xterm';
import 'xterm/css/xterm.css';

export default {
  name: 'TerminalComponent',
  data() {
    return {
      terminal: null, // xterm 实例
      socket: null,   // WebSocket 实例
    };
  },
  mounted() {
    this.initTerminal();
    this.initWebSocket();
  },
  methods: {
    // 初始化终端
    initTerminal() {
      this.terminal = new Terminal({
        rows: 30,
        cols: 80,
      });
      this.terminal.open(document.getElementById('terminal'));

      // 捕捉用户输入并发送到 WebSocket
      this.terminal.onData((data) => {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
          this.socket.send(data);
        }
      });
    },

    // 初始化 WebSocket 连接
    initWebSocket() {
      this.socket = new WebSocket('ws://localhost:1024/container/terminal');

      // 处理 WebSocket 消息，显示在终端
      this.socket.onmessage = (event) => {
        this.terminal.write(event.data);
      };

      // 处理 WebSocket 关闭
      this.socket.onclose = () => {
        this.terminal.write("\r\nConnection closed.");
      };

      // 处理 WebSocket 错误
      this.socket.onerror = (error) => {
        console.error("WebSocket error:", error);
        this.terminal.write("\r\nConnection error.");
      };
    },
  },
};
</script>

<style>
#terminal-container {
  width: 100%;
  height: 100%;
}

#terminal {
  width: 100%;
  height: 100%;
  background-color: black;
}
</style>
