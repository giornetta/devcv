<template>
    <div class="is-pulled-right">
        <div class="led" :class="{empty: level < 1}"></div>
        <div class="led" :class="{empty: level < 2}"></div>
        <div class="led" :class="{empty: level < 3}"></div>
        <div class="led" :class="{empty: level < 4}"></div>
    </div>
    
</template>

<script>
export default {
    props: {
        level: {
            type: Number,
            default: 0
        }
    }
}
</script>

<style scoped>
.led {
    display: inline-block;
    margin-top: 2px;
    height: 12px;
    width: 5px;
    margin-left: 5px;
    border-radius: 5px;
    background-color: hsl(217, 71%, 53%);
}

.empty {
    background-color: lightgrey;
}
</style>
