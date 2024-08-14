<script setup lang="ts">
import { cn } from '@/core/utils/tailwind.utils'
import {
  ProgressIndicator,
  ProgressRoot,
  type ProgressRootProps,
} from 'radix-vue'
import { computed, type HTMLAttributes } from 'vue'

const props = withDefaults(
  defineProps<ProgressRootProps & { class?: HTMLAttributes['class'], indicatorClass?: HTMLAttributes['class'] }>(),
  {
    modelValue: 0,
  },
)

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props

  return delegated
})
</script>

<template>
	<ProgressRoot
		v-bind="delegatedProps"
		:class="
      cn(
        'relative h-2 w-full overflow-hidden rounded-full bg-primary/20',
        props.class,
      )
    "
	>
		<ProgressIndicator
			class="h-full w-full flex-1 transition-all"
			:class="cn('bg-primary', props.indicatorClass)"
			:style="`transform: translateX(-${100 - (props.modelValue ?? 0)}%);`"
		/>
	</ProgressRoot>
</template>
