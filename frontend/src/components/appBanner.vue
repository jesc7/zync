<template>
    <svg height="40" width="360" xmlns="http://www.w3.org/2000/svg">
        <path d="M 120 26 h 120" id="myway1" fill="transparent" />
        <path d="M 250 6 h 80 a 10 10 0 0 1 0 26 h -300 a 10 10 0 0 1 0 -26 h 80" id="myway2" fill="transparent" />
        <text style="font: 13px Arial;">
            <textPath href="#myway1">
            <tspan id="big" style="font-size: 2.2rem; font-weight:bold; user-select: none;">
                ZFix UI 
            </tspan>
            </textPath>

            <textPath href="#myway2">
            <tspan id="small" style="font-weight:bold; font-size:0.46rem; letter-spacing: 2px;"> * makes you a bit happier * makes you a bit happier * makes you a bit happier * makes you a bit happier * </tspan>
            </textPath>
        </text>
    </svg>
</template>

<script setup>
</script>

<style scoped lang="scss">
#big {
    fill: $blue-8;
}
#small {
    fill: $grey-8;
}
</style>