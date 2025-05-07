<script lang="ts">
  import { createForm } from '@superform/client';
  import { z } from 'zod';
  import { toast } from 'svelte-sonner';
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  
  // Define the form schema using Zod
  const reportSchema = z.object({
    title: z.string().min(5, 'Title must be at least 5 characters'),
    description: z.string().min(20, 'Description must be at least 20 characters'),
    projectId: z.string().uuid('Invalid project ID'),
    departmentId: z.string().uuid('Invalid department ID'),
    securityScore: z.number().int().min(0).max(10),
    performanceScore: z.number().int().min(0).max(10),
    memoryScore: z.number().int().min(0).max(10),
    testingScore: z.number().int().min(0).max(10),
    errorScore: z.number().int().min(0).max(10),
    loadScore: z.number().int().min(0).max(10),
    securityDetails: z.string().optional(),
    performanceDetails: z.string().optional(),
    memoryDetails: z.string().optional(),
    testingDetails: z.string().optional(),
    errorDetails: z.string().optional(),
    loadDetails: z.string().optional(),
    tags: z.array(z.string()).optional()
  });
  
  // Create the form
  const { form, errors, enhance, constraints } = createForm(reportSchema, {
    onSubmit: async ({ formData, cancel }) => {
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      // Check if we want to submit or save as draft
      const submitType = formData.get('submitType');
      
      if (submitType === 'draft') {
        toast.success('Report saved as draft');
      } else {
        toast.success('Report submitted successfully');
      }
    },
    onError: ({ result }) => {
      toast.error('Failed to submit report');
    }
  });
  
  // Mock data for dropdowns
  const projects = [
    { id: '123e4567-e89b-12d3-a456-426614174000', name: 'E-commerce Platform' },
    { id: '223e4567-e89b-12d3-a456-426614174000', name: 'CRM System' },
    { id: '323e4567-e89b-12d3-a456-426614174000', name: 'Mobile Banking App' },
    { id: '423e4567-e89b-12d3-a456-426614174000', name: 'Content Management System' }
  ];
  
  const departments = [
    { id: '523e4567-e89b-12d3-a456-426614174000', name: 'Security' },
    { id: '623e4567-e89b-12d3-a456-426614174000', name: 'Backend' },
    { id: '723e4567-e89b-12d3-a456-426614174000', name: 'Frontend' },
    { id: '823e4567-e89b-12d3-a456-426614174000', name: 'UI/UX' },
    { id: '923e4567-e89b-12d3-a456-426614174000', name: 'Database' },
    { id: 'a23e4567-e89b-12d3-a456-426614174000', name: 'DevOps' }
  ];
  
  // Animation references
  let formRef: HTMLElement;
  let scoreRef: HTMLElement;
  
  onMount(() => {
    // Animate form sections
    gsap.from(formRef.querySelectorAll('.form-section'), {
      y: 20,
      opacity: 0,
      duration: 0.6,
      stagger: 0.1,
      ease: 'power2.out'
    });
    
    // Animate score sliders
    gsap.from(scoreRef.querySelectorAll('.score-slider'), {
      width: 0,
      opacity: 0,
      duration: 0.8,
      stagger: 0.1,
      delay: 0.4,
      ease: 'power2.out'
    });
  });
  
  // Helper function to get color based on score
  function getScoreColor(score: number): string {
    if (score >= 8) return 'bg-green-500';
    if (score >= 6) return 'bg-amber-500';
    if (score >= 4) return 'bg-orange-500';
    return 'bg-red-500';
  }
  
  // Handle form submission
  function handleSubmit(event: Event, type: 'draft' | 'submit') {
    const formData = new FormData(event.target as HTMLFormElement);
    formData.append('submitType', type);
  }
</script>

<div bind:this={formRef} class="space-y-8">
  <div class="form-section space-y-4">
    <h2 class="text-xl font-semibold">Project Information</h2>
    
    <div class="grid gap-4 md:grid-cols-2">
      <div class="space-y-2">
        <label for="title" class="text-sm font-medium">Title</label>
        <input
          id="title"
          name="title"
          bind:value={$form.title}
          class="w-full rounded-md border border-input bg-background px-3 py-2"
          placeholder="Enter report title"
        />
        {#if $errors.title}
          <p class="text-sm text-red-500">{$errors.title}</p>
        {/if}
      </div>
      
      <div class="space-y-2">
        <label for="projectId" class="text-sm font-medium">Project</label>
        <select
          id="projectId"
          name="projectId"
          bind:value={$form.projectId}
          class="w-full rounded-md border border-input bg-background px-3 py-2"
        >
          <option value="">Select a project</option>
          {#each projects as project}
            <option value={project.id}>{project.name}</option>
          {/each}
        </select>
        {#if $errors.projectId}
          <p class="text-sm text-red-500">{$errors.projectId}</p>
        {/if}
      </div>
    </div>
    
    <div class="space-y-2">
      <label for="description" class="text-sm font-medium">Description</label>
      <textarea
        id="description"
        name="description"
        bind:value={$form.description}
        class="w-full rounded-md border border-input bg-background px-3 py-2 min-h-[100px]"
        placeholder="Provide a detailed description of the project report"
      ></textarea>
      {#if $errors.description}
        <p class="text-sm text-red-500">{$errors.description}</p>
      {/if}
    </div>
    
    <div class="space-y-2">
      <label for="departmentId" class="text-sm font-medium">Department</label>
      <select
        id="departmentId"
        name="departmentId"
        bind:value={$form.departmentId}
        class="w-full rounded-md border border-input bg-background px-3 py-2"
      >
        <option value="">Select a department</option>
        {#each departments as department}
          <option value={department.id}>{department.name}</option>
        {/each}
      </select>
      {#if $errors.departmentId}
        <p class="text-sm text-red-500">{$errors.departmentId}</p>
      {/if}
    </div>
  </div>
  
  <div bind:this={scoreRef} class="form-section space-y-6">
    <h2 class="text-xl font-semibold">Evaluation Scores</h2>
    <p class="text-sm text-muted-foreground">Rate each component on a scale of 0-10.</p>
    
    <div class="space-y-4">
      <!-- Security Score -->
      <div class="space-y-2">
        <div class="flex items-center justify-between">
          <label for="securityScore" class="text-sm font-medium">Security (S)</label>
          <span class="text-sm font-medium">{$form.securityScore || 0}/10</span>
        </div>
        <div class="flex items-center gap-2">
          <input
            id="securityScore"
            name="securityScore"
            type="range"
            min="0"
            max="10"
            step="1"
            bind:value={$form.securityScore}
            class="w-full"
          />
        </div>
        <div class="h-2 w-full bg-gray-200 rounded-full overflow-hidden">
          <div class="score-slider h-full {getScoreColor($form.securityScore || 0)}" style="width: {($form.securityScore || 0) * 10}%"></div>
        </div>
        <textarea
          name="securityDetails"
          bind:value={$form.securityDetails}
          class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
          placeholder="Provide details about security implementation"
          rows="2"
        ></textarea>
      </div>
      
      <!-- Performance Score -->
      <div class="space-y-2">
        <div class="flex items-center justify-between">
          <label for="performanceScore" class="text-sm font-medium">Performance (P)</label>
          <span class="text-sm font-medium">{$form.performanceScore || 0}/10</span>
        </div>
        <div class="flex items-center gap-2">
          <input
            id="performanceScore"
            name="performanceScore"
            type="range"
            min="0"
            max="10"
            step="1"
            bind:value={$form.performanceScore}
            class="w-full"
          />
        </div>
        <div class="h-2 w-full bg-gray-200 rounded-full overflow-hidden">
          <div class="score-slider h-full {getScoreColor($form.performanceScore || 0)}" style="width: {($form.performanceScore || 0) * 10}%"></div>
        </div>
        <textarea
          name="performanceDetails"
          bind:value={$form.performanceDetails}
          class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
          placeholder="Provide details about performance optimizations"
          rows="2"
        ></textarea>
      </div>
      
      <!-- Memory Score -->
      <div class="space-y-2">
        <div class="flex items-center justify-between">
          <label for="memoryScore" class="text-sm font-medium">Memory (M)</label>
          <span class="text-sm font-medium">{$form.memoryScore || 0}/10</span>
        </div>
        <div class="flex items-center gap-2">
          <input
            id="memoryScore"
            name="memoryScore"
            type="range"
            min="0"
            max="10"
            step="1"
            bind:value={$form.memoryScore}
            class="w-full"
          />
        </div>
        <div class="h-2 w-full bg-gray-200 rounded-full overflow-hidden">
          <div class="score-slider h-full {getScoreColor($form.memoryScore || 0)}" style="width: {($form.memoryScore || 0) * 10}%"></div>
        </div>
        <textarea
          name="memoryDetails"
          bind:value={$form.memoryDetails}
          class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
          placeholder="Provide details about memory management"
          rows="2"
        ></textarea>
      </div>
      
      <!-- Testing Score -->
      <div class="space-y-2">
        <div class="flex items-center justify-between">
          <label for="testingScore" class="text-sm font-medium">Testing (T)</label>
          <span class="text-sm font-medium">{$form.testingScore || 0}/10</span>
        </div>
        <div class="flex items-center gap-2">
          <input
            id="testingScore"
            name="testingScore"
            type="range"
            min="0"
            max="10"
            step="1"
            bind:value={$form.testingScore}
            class="w-full"
          />
        </div>
        <div class="h-2 w-full bg-gray-200 rounded-full overflow-hidden">
          <div class="score-slider h-full {getScoreColor($form.testingScore || 0)}" style="width: {($form.testingScore || 0) * 10}%"></div>
        </div>
        <textarea
          name="testingDetails"
          bind:value={$form.testingDetails}
          class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
          placeholder="Provide details about testing coverage"
          rows="2"
        ></textarea>
      </div>
      
      <!-- Error Score -->
      <div class="space-y-2">
        <div class="flex items-center justify-between">
          <label for="errorScore" class="text-sm font-medium">Error Handling (E)</label>
          <span class="text-sm font-medium">{$form.errorScore || 0}/10</span>
        </div>
        <div class="flex items-center gap-2">
          <input
            id="errorScore"
            name="errorScore"
            type="range"
            min="0"
            max="10"
            step="1"
            bind:value={$form.errorScore}
            class="w-full"
          />
        </div>
        <div class="h-2 w-full bg-gray-200 rounded-full overflow-hidden">
          <div class="score-slider h-full {getScoreColor($form.errorScore || 0)}" style="width: {($form.errorScore || 0) * 10}%"></div>
        </div>
        <textarea
          name="errorDetails"
          bind:value={$form.errorDetails}
          class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
          placeholder="Provide details about error handling"
          rows="2"
        ></textarea>
      </div>
      
      <!-- Load Score -->
      <div class="space-y-2">
        <div class="flex items-center justify-between">
          <label for="loadScore" class="text-sm font-medium">Load Impact (L)</label>
          <span class="text-sm font-medium">{$form.loadScore || 0}/10</span>
        </div>
        <div class="flex items-center gap-2">
          <input
            id="loadScore"
            name="loadScore"
            type="range"
            min="0"
            max="10"
            step="1"
            bind:value={$form.loadScore}
            class="w-full"
          />
        </div>
        <div class="h-2 w-full bg-gray-200 rounded-full overflow-hidden">
          <div class="score-slider h-full {getScoreColor($form.loadScore || 0)}" style="width: {($form.loadScore || 0) * 10}%"></div>
        </div>
        <textarea
          name="loadDetails"
          bind:value={$form.loadDetails}
          class="w-full rounded-md border border-input bg-background px-3 py-2 text-sm"
          placeholder="Provide details about load impact"
          rows="2"
        ></textarea>
      </div>
    </div>
  </div>
  
  <div class="form-section">
    <div class="flex items-center justify-between">
      <div class="text-sm text-muted-foreground">
        <div class="flex items-center gap-1">
          <span class="font-medium">Overall Score:</span>
          <span class="font-bold {getScoreColor(
            (($form.securityScore || 0) + 
            ($form.performanceScore || 0) + 
            ($form.memoryScore || 0) + 
            ($form.testingScore || 0) + 
            ($form.errorScore || 0) + 
            ($form.loadScore || 0)) / 6
          )}">
            {(
              (($form.securityScore || 0) + 
              ($form.performanceScore || 0) + 
              ($form.memoryScore || 0) + 
              ($form.testingScore || 0) + 
              ($form.errorScore || 0) + 
              ($form.loadScore || 0)) / 6
            ).toFixed(1)}/10
          </span>
        </div>
        <div class="mt-1">
          Score: [S{$form.securityScore || 0},P{$form.performanceScore || 0},M{$form.memoryScore || 0},T{$form.testingScore || 0},E{$form.errorScore || 0},L{$form.loadScore || 0}]
        </div>
      </div>
      
      <div class="flex gap-2">
        <button
          type="button"
          class="rounded-md bg-muted px-4 py-2 text-sm font-medium text-foreground hover:bg-muted/80"
          on:click={(e) => handleSubmit(e, 'draft')}
        >
          Save as Draft
        </button>
        <button
          type="submit"
          class="rounded-md bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"
          on:click={(e) => handleSubmit(e, 'submit')}
        >
          Submit Report
        </button>
      </div>
    </div>
  </div>
</div>
