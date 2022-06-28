<script lang="ts">
  import { onMount } from "svelte";
  import QuoteDisplay from "./lib/QuoteDisplay.svelte";
  import Spinner from "./lib/Spinner.svelte";
  import type Quote from "./model/quote";

  let quote: Quote | null = null;
  let isLoading: boolean = false;

  async function update() {
    isLoading = true;
    const response = await fetch("/api/quotes/random");
    quote = (await response.json()) as Quote;
    isLoading = false;
  }

  onMount(() => {
    update();
  });
</script>

<main>
  <h1>Your random quote is:</h1>
  {#if isLoading}
    <Spinner />
  {:else}
    <QuoteDisplay {quote} />
    <button on:click={update}>Next quote</button>
  {/if}
</main>

<style>
  :root {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
      Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  }

  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    max-width: 800px;
    margin: auto;
  }

  h1 {
    font-family: "Courier New", Courier, monospace;
    font-size: 2.8rem;
    font-weight: bolder;
    text-transform: uppercase;
  }

  button {
    border: 2px solid black;
    border-radius: 5px;
    background-color: white;
    font-weight: bold;
    text-transform: uppercase;
    font-size: 1.2em;
    margin-top: 3em;
  }
  button:hover {
    cursor: pointer;
    filter: invert();
  }
</style>
