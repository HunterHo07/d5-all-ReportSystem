<script lang="ts">
  import ReportForm from '$lib/components/ReportForm.svelte';
</script>

<div class="container py-8">
  <div class="mb-8">
    <h1 class="text-3xl font-bold tracking-tight">Create New Report</h1>
    <p class="text-muted-foreground">Fill out the form below to create a new project report.</p>
  </div>
  
  <div class="rounded-lg border bg-card shadow-sm">
    <div class="p-6">
      <ReportForm />
    </div>
  </div>
</div>
