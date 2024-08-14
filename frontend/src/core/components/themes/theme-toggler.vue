<script setup lang="ts">
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from '@/core/components/ui/dropdown-menu'
import { useColorMode } from '@vueuse/core'
import { computed } from 'vue'

const mode = useColorMode()

const iconsByMode = {
  'light': "pi pi-moon",
  'dark': "pi pi-sun",
  'auto': "pi pi-palette"
}

const iconType = computed(() => iconsByMode[mode.value])
</script>

<template>
	<DropdownMenu>
		<DropdownMenuTrigger as-child>
			<div variant="outline">
				<span
					class="text-lg"
					:class="iconType"
				/>
				<span class="sr-only">Toggle theme</span>
			</div>
		</DropdownMenuTrigger>
		<DropdownMenuContent align="end">
			<DropdownMenuItem @click="mode = 'light'"> Light </DropdownMenuItem>
			<DropdownMenuItem @click="mode = 'dark'"> Dark </DropdownMenuItem>
			<DropdownMenuItem @click="mode = 'auto'"> System </DropdownMenuItem>
		</DropdownMenuContent>
	</DropdownMenu>
</template>
