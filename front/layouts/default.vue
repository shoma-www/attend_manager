<template>
  <div class="sans-serif fs-normal">
    <header class="mb4">
      <div class="bg-black-80">
        <nav class="dt w-100 pa3 center">
          <div class="dtc v-mid">
            <div class="f8">
              <a class="hover-white no-underline white-90" href="/">
                勤怠管理野郎
              </a>
            </div>
          </div>
          <div class="dtc v-mid tr">
            <a
              class="f6 fw4 hover-white no-underline white-70 dib ml2 pv2 ph3 ba"
              href="/signin"
            >
              SignIn
            </a>
            <a
              class="f6 fw4 hover-white no-underline white-70 dib ml2 pv2 ph3 ba"
              href="/signup"
            >
              SignUp
            </a>
          </div>
        </nav>
      </div>
    </header>
    <div class="w-90 mw8 center">
      <Nuxt />
    </div>
  </div>
</template>
