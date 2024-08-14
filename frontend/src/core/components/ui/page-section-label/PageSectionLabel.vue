<script setup lang="ts">
import { cn } from '@/core/utils/tailwind.utils'

import { computed, type HTMLAttributes } from 'vue'

const props = withDefaults(
  defineProps<{ class?: HTMLAttributes['class']}>(),
	{}
)

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props

  return delegated
})
</script>

<template>
	<h3
		v-bind="delegatedProps"
		:class="
      cn(
        'text-lg font-medium',
        props.class,
      )
    "
	>
		<slot></slot>
	</h3>
</template>
