<div class="lds-dual-ring" />

<style>
  .lds-dual-ring {
    margin-top: 3em;
    width: 80px;
    height: 80px;
  }
  .lds-dual-ring:after {
    content: " ";
    display: block;
    width: 64px;
    height: 64px;
    margin: 8px;
    border-radius: 50%;
    border: 6px solid hsl(0, 0%, 5%);
    border-color: hsl(0, 0%, 5%) transparent hsl(0, 0%, 5%) transparent;
    animation: lds-dual-ring 2s linear infinite;
  }
  @keyframes lds-dual-ring {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>
