<template>
    <div>
        {{champion}}
    </div>
</template>

<script setup lang="ts">
import { useDataDragonStore } from '@/stores/DataDragonStore';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

    // Load aatrox builds
    const route = useRoute()

    const props = defineProps({
        champion: String
    })

    const championKey = props.champion || "";

    const dataDragonStore = useDataDragonStore()

    const champion = computed(() => {
       console.log("championKey:", championKey)
       return dataDragonStore.champions.get(championKey)
    })
    
</script>