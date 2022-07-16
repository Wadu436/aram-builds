<template>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    style="display: block; shape-rendering: auto"
    width="4rem"
    viewBox="0 0 100 100"
    preserveAspectRatio="xMidYMid"
  >
    <circle
      cx="50"
      cy="50"
      fill="none"
      stroke="currentColor"
      stroke-width="10"
      r="35"
      stroke-dasharray="185%"
    >
      <animate
        attributeName="stroke-dasharray"
        values="185%;140%;185%"
        dur="1s"
        repeatCount="indefinite"
      ></animate>
    </circle>
  </svg>
</template>
