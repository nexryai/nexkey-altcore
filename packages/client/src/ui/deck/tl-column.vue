<template>
<XColumn :menu="menu" :column="column" :is-stacked="isStacked" :indicated="indicated" @change-active-state="onChangeActiveState" @parent-focus="$event => emit('parent-focus', $event)">
    <template #header>
        <i v-if="column.tl === 'home'" class="ti ti-home"></i>
        <i v-else-if="column.tl === 'social'" class="ti ti-rocket"></i>
        <i v-else-if="column.tl === 'global'" class="ti ti-whirl"></i>
        <span style="margin-left: 8px;">{{ column.name }}</span>
    </template>

    <div v-if="disabled || ((column.tl === 'local' || column.tl === 'social') && !enableLTL) || (column.tl === 'media' && (!enableMTL || !enableLTL)) || (column.tl === 'global' && !enableGTL) || (column.tl === 'personal' && !enablePTL) || (column.tl === 'limited' && !enableLimitedTL)" class="iwaalbte">
        <p>
            <i class="ti ti-minus-circle"></i>
            {{ i18n.ts.disabledTimelineTitle }}
        </p>
        <p class="desc">{{ i18n.ts.disabledTimelineDescription }}</p>
    </div>
    <XTimeline v-else-if="column.tl" ref="timeline" :key="column.tl" :src="column.tl" @after="() => emit('loaded')" @queue="queueUpdated" @note="onNote"/>
</XColumn>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import XColumn from "./column.vue";
import { removeColumn, updateColumn, Column } from "./deck-store";
import XTimeline from "@/components/MkTimeline.vue";
import * as os from "@/os";
import { $i } from "@/account";
import { instance } from "@/instance";
import { i18n } from "@/i18n";
import { defaultStore } from "@/store";

const props = defineProps<{
	column: Column;
	isStacked: boolean;
}>();

const emit = defineEmits<{
	(ev: "loaded"): void;
	(ev: "parent-focus", direction: "up" | "down" | "left" | "right"): void;
}>();

let disabled = ref(false);
let indicated = ref(false);
let columnActive = ref(true);
let enableMTL = ref(false);
let enableLTL = ref(false);
let enableGTL = ref(false);
let enablePTL = ref(false);
let enableLimitedTL = ref(false);

onMounted(() => {
    if (props.column.tl == null) {
        setType();
    } else if ($i) {
        disabled.value = !$i.isModerator && !$i.isAdmin && (
            instance.disableLocalTimeline && ["local", "social"].includes(props.column.tl) ||
			instance.disableGlobalTimeline && ["global"].includes(props.column.tl));
    }
    enableLTL.value = defaultStore.state.enableLTL;
    enableLimitedTL.value = defaultStore.state.enableLimitedTL;
    enableMTL.value = defaultStore.state.enableMTL;
    enableGTL.value = defaultStore.state.enableGTL;
    enablePTL.value = defaultStore.state.enablePTL;
});

async function setType() {
    const { canceled, result: src } = await os.select({
        title: i18n.ts.timeline,
        items: [{
            value: "home" as const, text: i18n.ts._timelines.home,
        }, {
            value: "local" as const, text: i18n.ts._timelines.local,
        }, {
            value: "social" as const, text: i18n.ts._timelines.social,
        }, {
            value: "global" as const, text: i18n.ts._timelines.global,
        }],
    });
    if (canceled) {
        if (props.column.tl == null) {
            removeColumn(props.column.id);
        }
        return;
    }
    updateColumn(props.column.id, {
        tl: src,
    });
}

function queueUpdated(q) {
    if (columnActive.value) {
        indicated.value = q !== 0;
    }
}

function onNote() {
    if (!columnActive.value) {
        indicated.value = true;
    }
}

function onChangeActiveState(state) {
    columnActive.value = state;

    if (columnActive.value) {
        indicated.value = false;
    }
}

const menu = [{
    icon: "ti ti-pencil",
    text: i18n.ts.timeline,
    action: setType,
}];
</script>

<style lang="scss" scoped>
.iwaalbte {
	text-align: center;

	> p {
		margin: 16px;

		&.desc {
			font-size: 14px;
		}
	}
}
</style>
