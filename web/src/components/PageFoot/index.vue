<template>
  <footer>
    <div class="wrapper">
      <div>
        <a href="https://github.com/DokiDoki1103/tikuAdapter" class="span-tooltip" target="_blank">
          <i>
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-github"
              viewBox="0 0 16 16">
              <path
                d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27s1.36.09 2 .27c1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.01 8.01 0 0 0 16 8c0-4.42-3.58-8-8-8">
              </path>
            </svg>
          </i>
          <span>Github</span>
        </a>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'PageFoot',
}
</script>

<style scoped>
footer {
  height: 60px;
  width: 100%;
  background-color: #fff;
  display: flex;
  align-items: center;
  justify-content: right;
  padding-right: 3%;
}

.wrapper {
  display: flex;
  gap: 10px;
  font-family: monospace;

}

.wrapper>div {
  text-align: center;
  
}

.wrapper>div+div {
  border-inline-start: 1px solid #d6cfcf;
  padding-inline-start: 15px;
}

.wrapper>div>p {
  margin-top: 15px;
  max-width: 200px;
}

.wrapper a {
  display: inline-flex;
  width: 30px;
  height: 30px;
  border-radius: 100%;
  justify-content: center;
  align-items: center;
  background-color: #191919;
  color: #d5ad72;
  position: relative;
}

.data-tooltip::before,
.data-tooltip::after {
  position: absolute;
  visibility: hidden;
  opacity: 0;
  transform: translateY(5px);
  transition: all 0.2s ease;
}

.data-tooltip::before {
  content: attr(data-tooltip-content);
  width: fit-content;
  font-size: small;
  font-weight: 600;
  background-color: inherit;
  padding: 3px 10px;
  border-radius: 4px;
  bottom: calc(100% + 5px);
}

.data-tooltip::after {
  content: "";
  border-top: 5px solid #191919;
  border-left: 5px solid transparent;
  border-right: 5px solid transparent;
  bottom: calc(100% + 1px);
}

.data-tooltip:hover::before,
.data-tooltip:hover::after {
  visibility: visible;
  opacity: 1;
  transform: translateY(0);
}

.span-tooltip>i {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.span-tooltip>span {
  position: absolute;
  width: fit-content;
  font-size: small;
  font-weight: 600;
  background-color: inherit;
  padding: 3px 10px;
  border-radius: 4px;
  bottom: 100%;
  visibility: hidden;
  opacity: 0;
  transition: all 0.2s ease;
}

.span-tooltip>span::before {
  position: absolute;
  content: "";
  width: 10px;
  height: 10px;
  border-radius: 3px;
  background-color: inherit;
  transform: rotate(45deg);
  top: calc(100% - 6px);
  left: 40%;
}

.span-tooltip>i:hover+span {
  visibility: visible;
  opacity: 1;
  bottom: calc(100% + 6px);
}
</style>
