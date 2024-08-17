<template>
	<div v-if="!selectedCategoryId">
		<HabitsCreateCategoriesView @select="confirmCategorySelection" />
	</div>

	<div v-if="selectedCategoryId">
		<HabitsCreateFormView @create="onCreate" />
	</div>
</template>

<script setup lang="ts">
import { ROUTES } from '@/core/config/routes.config'
import HabitsCreateCategoriesView from '@/modules/habits/components/views/habits-create-categories-view.vue'
import HabitsCreateFormView from '@/modules/habits/components/views/habits-create-form-view.vue'
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const route = useRoute()

const selectedCategoryId = computed(() => Number(route.query.categoryId))

const confirmCategorySelection = (categoryId: number | null) => {
  if (categoryId !== null) {
    router.push({ query: { categoryId: categoryId }})
  }
}

const onCreate = () => {
	router.push({ name: ROUTES.HOME})

}
</script>

<style scoped></style>
