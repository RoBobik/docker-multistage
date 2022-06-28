<script lang="ts">
  import type Quote from "../model/quote";

  export let quote: Quote;
</script>

<div>
  <p class="text">{quote?.text}</p>
  <p class="author">{quote?.author}</p>
</div>

<style>
  p {
    margin: 0.25em;
  }
  p.text {
    color: hsl(210, 10%, 10%);
    font-family: cursive;
    font-size: 2.5em;
  }
  p.text::before {
    content: "“";
  }
  p.text::after {
    content: "”";
  }
  p.author {
    font-family: sans-serif;
    font-size: 1.25em;
    text-align: right;
    text-transform: uppercase;
    padding-right: 1em;
  }
  p.author::before {
    content: "-- ";
  }
</style>
