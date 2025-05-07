<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  
  // Mock data for dashboard
  const recentReports = [
    { id: 1, title: 'Frontend Authentication Module', status: 'submitted', department: 'Security', date: '2023-11-15', score: { s: 8, p: 7, m: 6, t: 7, e: 8, l: 6 } },
    { id: 2, title: 'API Gateway Implementation', status: 'draft', department: 'Backend', date: '2023-11-14', score: { s: 9, p: 8, m: 7, t: 8, e: 7, l: 8 } },
    { id: 3, title: 'User Dashboard Redesign', status: 'approved', department: 'UI/UX', date: '2023-11-10', score: { s: 7, p: 9, m: 8, t: 6, e: 8, l: 7 } },
    { id: 4, title: 'Database Optimization', status: 'rejected', department: 'Database', date: '2023-11-08', score: { s: 6, p: 8, m: 9, t: 7, e: 6, l: 9 } }
  ];
  
  const departments = [
    { id: 1, name: 'Security', count: 12 },
    { id: 2, name: 'Backend', count: 24 },
    { id: 3, name: 'Frontend', count: 18 },
    { id: 4, name: 'UI/UX', count: 15 },
    { id: 5, name: 'Database', count: 9 },
    { id: 6, name: 'DevOps', count: 7 }
  ];
  
  const stats = [
    { label: 'Total Reports', value: 128 },
    { label: 'Pending Approval', value: 23 },
    { label: 'Approved', value: 86 },
    { label: 'Average Score', value: '7.4/10' }
  ];
  
  // Animation references
  let statsRef: HTMLElement;
  let reportsRef: HTMLElement;
  
  onMount(() => {
    // Animate stats cards
    gsap.from(statsRef.querySelectorAll('.stat-card'), {
      y: 20,
      opacity: 0,
      duration: 0.6,
      stagger: 0.1,
      ease: 'power2.out'
    });
    
    // Animate reports table
    gsap.from(reportsRef.querySelectorAll('tr'), {
      y: 10,
      opacity: 0,
      duration: 0.4,
      stagger: 0.05,
      delay: 0.3,
      ease: 'power1.out'
    });
  });
  
  // Helper function to get status color
  function getStatusColor(status: string): string {
    switch (status) {
      case 'submitted': return 'bg-blue-100 text-blue-800';
      case 'draft': return 'bg-gray-100 text-gray-800';
      case 'approved': return 'bg-green-100 text-green-800';
      case 'rejected': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  }
  
  // Helper function to get score color
  function getScoreColor(score: number): string {
    if (score >= 8) return 'text-green-600';
    if (score >= 6) return 'text-amber-600';
    return 'text-red-600';
  }
</script>

<div class="container py-8">
  <div class="flex items-center justify-between mb-8">
    <div>
      <h1 class="text-3xl font-bold tracking-tight">Dashboard</h1>
      <p class="text-muted-foreground">Welcome back! Here's an overview of your project reports.</p>
    </div>
    <a href="/reports/new" class="inline-flex items-center justify-center rounded-md bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90">
      New Report
    </a>
  </div>
  
  <!-- Stats Overview -->
  <div bind:this={statsRef} class="grid gap-4 md:grid-cols-2 lg:grid-cols-4 mb-8">
    {#each stats as stat}
      <div class="stat-card rounded-lg border bg-card p-6 shadow-sm">
        <div class="text-sm font-medium text-muted-foreground">{stat.label}</div>
        <div class="text-3xl font-bold">{stat.value}</div>
      </div>
    {/each}
  </div>
  
  <div class="grid gap-8 md:grid-cols-7">
    <!-- Recent Reports -->
    <div class="md:col-span-5">
      <div class="rounded-lg border bg-card shadow-sm">
        <div class="p-6">
          <h2 class="text-xl font-semibold">Recent Reports</h2>
          <p class="text-sm text-muted-foreground">Your most recent project reports and their status.</p>
        </div>
        <div class="px-6 pb-6">
          <div class="rounded-md border">
            <div bind:this={reportsRef} class="relative w-full overflow-auto">
              <table class="w-full caption-bottom text-sm">
                <thead>
                  <tr class="border-b transition-colors hover:bg-muted/50">
                    <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground">Title</th>
                    <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground">Status</th>
                    <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground">Department</th>
                    <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground">Date</th>
                    <th class="h-12 px-4 text-left align-middle font-medium text-muted-foreground">Score</th>
                  </tr>
                </thead>
                <tbody>
                  {#each recentReports as report}
                    <tr class="border-b transition-colors hover:bg-muted/50">
                      <td class="p-4 align-middle font-medium">{report.title}</td>
                      <td class="p-4 align-middle">
                        <span class={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${getStatusColor(report.status)}`}>
                          {report.status}
                        </span>
                      </td>
                      <td class="p-4 align-middle">{report.department}</td>
                      <td class="p-4 align-middle">{report.date}</td>
                      <td class="p-4 align-middle">
                        <div class="flex items-center gap-1">
                          <span class={getScoreColor(report.score.s)}>S{report.score.s}</span>
                          <span class={getScoreColor(report.score.p)}>P{report.score.p}</span>
                          <span class={getScoreColor(report.score.m)}>M{report.score.m}</span>
                          <span class={getScoreColor(report.score.t)}>T{report.score.t}</span>
                          <span class={getScoreColor(report.score.e)}>E{report.score.e}</span>
                          <span class={getScoreColor(report.score.l)}>L{report.score.l}</span>
                        </div>
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>
          </div>
          <div class="mt-4 flex justify-end">
            <a href="/reports" class="text-sm text-primary hover:underline">View all reports →</a>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Departments -->
    <div class="md:col-span-2">
      <div class="rounded-lg border bg-card shadow-sm h-full">
        <div class="p-6">
          <h2 class="text-xl font-semibold">Departments</h2>
          <p class="text-sm text-muted-foreground">Reports by department.</p>
        </div>
        <div class="px-6 pb-6">
          <ul class="space-y-2">
            {#each departments as dept}
              <li class="flex items-center justify-between rounded-md border px-4 py-3 text-sm">
                <span>{dept.name}</span>
                <span class="rounded-full bg-primary/10 px-2 py-1 text-xs font-medium text-primary">{dept.count}</span>
              </li>
            {/each}
          </ul>
        </div>
      </div>
    </div>
  </div>
</div>
