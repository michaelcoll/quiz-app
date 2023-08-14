<script setup lang="ts">
  const props = defineProps<{
    remainingSec?: number | null;
    quizDuration?: number | null;
  }>();

  const internalRemainingSec = ref<number>();
  const remainingTime = ref<string>();

  onMounted(() => {
    if (props.remainingSec !== undefined && props.remainingSec != null) {
      internalRemainingSec.value = props.remainingSec.valueOf();
      displayTimeLeft(props.remainingSec);
      countdown(props.remainingSec);
    }
  });

  function countdown(sec: number) {
    internalRemainingSec.value = sec;

    const intervalTimer = setInterval(() => {
      internalRemainingSec.value = internalRemainingSec.value! - 1;

      if (internalRemainingSec.value < 0) {
        clearInterval(intervalTimer);
        return;
      }
      displayTimeLeft(internalRemainingSec.value);
    }, 1000);
  }

  function displayTimeLeft(sec: number) {
    const minutes = Math.floor((sec % 3600) / 60);
    const seconds = sec % 60;

    remainingTime.value = `${zeroPadded(minutes)}:${zeroPadded(seconds)}`;
  }

  function zeroPadded(num: number): string {
    // 4 --> 04
    return num < 10 ? `0${num}` : `${num}`;
  }
</script>

<template>
  <div
    v-if="props.quizDuration != undefined && props.remainingSec != undefined"
    class="relative w-56 pt-1">
    <div class="mb-2 flex items-center justify-between">
      <div>
        <span
          v-if="internalRemainingSec && internalRemainingSec > 0"
          class="inline-block rounded-full bg-green-200 px-2 py-1 text-xs font-semibold uppercase text-green-600">
          Session in progress
        </span>
        <span
          v-else
          class="inline-block rounded-full bg-gray-200 px-2 py-1 text-xs font-semibold uppercase text-gray-600">
          Session is over
        </span>
      </div>
      <div class="text-right">
        <span class="inline-block text-2xl font-semibold text-green-600">
          {{ remainingTime }}
        </span>
      </div>
    </div>
    <div class="mb-4 flex h-2 overflow-hidden rounded bg-green-200 text-xs">
      <div
        :style="`width: ${
          ((props.quizDuration - (internalRemainingSec ?? 0)) * 100) / props.quizDuration
        }%`"
        class="flex flex-col justify-center whitespace-nowrap bg-green-500 text-center text-white shadow-none"></div>
    </div>
  </div>
</template>

<style scoped></style>
