<script lang="ts">
  import '../app.css';
  import { page } from '$app/stores';
  import { Toaster } from 'svelte-sonner';
</script>

<div class="min-h-screen bg-background font-sans antialiased">
  <div class="relative flex min-h-screen flex-col">
    <header class="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div class="container flex h-16 items-center">
        <div class="flex items-center gap-2 font-bold">
          <div class="h-8 w-8 rounded-full bg-primary flex items-center justify-center text-white">PR</div>
          <span>ProjectReport</span>
        </div>
        <nav class="ml-auto flex gap-6">
          <a href="/" class={$page.url.pathname === '/' ? 'text-primary' : 'text-foreground hover:text-primary'}>Dashboard</a>
          <a href="/reports" class={$page.url.pathname.startsWith('/reports') ? 'text-primary' : 'text-foreground hover:text-primary'}>Reports</a>
          <a href="/analytics" class={$page.url.pathname.startsWith('/analytics') ? 'text-primary' : 'text-foreground hover:text-primary'}>Analytics</a>
          <a href="/settings" class={$page.url.pathname.startsWith('/settings') ? 'text-primary' : 'text-foreground hover:text-primary'}>Settings</a>
        </nav>
        <div class="ml-6 flex items-center gap-2">
          <button class="rounded-full h-8 w-8 bg-muted flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-bell"><path d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9"/><path d="M10.3 21a1.94 1.94 0 0 0 3.4 0"/></svg>
          </button>
          <div class="h-8 w-8 rounded-full bg-primary flex items-center justify-center text-white">JD</div>
        </div>
      </div>
    </header>
    <main class="flex-1">
      <slot />
    </main>
    <footer class="border-t py-6 md:py-0">
      <div class="container flex flex-col items-center justify-between gap-4 md:h-16 md:flex-row">
        <p class="text-sm text-muted-foreground">
          &copy; 2023 ProjectReport. All rights reserved.
        </p>
        <div class="flex gap-4 text-sm text-muted-foreground">
          <a href="/terms" class="hover:underline">Terms</a>
          <a href="/privacy" class="hover:underline">Privacy</a>
          <a href="/contact" class="hover:underline">Contact</a>
        </div>
      </div>
    </footer>
  </div>
</div>

<Toaster />
