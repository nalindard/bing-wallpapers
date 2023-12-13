<template>
    <div>
        <input type="radio" name="accent" class="accent-emerald-400 h-7 content-none w-7 bg-transparent">
        <input type="radio" name="accent" class="accent-emerald-400 h-7 content-none w-7 bg-transparent">
        <input type="radio" name="accent" class="accent-emerald-400 h-7 content-none w-7 bg-transparent">
        <input type="radio" name="accent" class="accent-emerald-400 h-7 content-none w-7 bg-transparent">
        <input type="radio" name="accent" class="accent-emerald-400 h-7 content-none w-7 bg-transparent">
        <input type="radio" name="accent" class="accent-emerald-400 h-7 content-none w-7 bg-transparent">
        <input type="radio" name="accent" class="accent-emerald-400 h-7 content-none w-7 bg-transparent">
    </div>
</template>

<style scoped>
input[type="radio"] {
	appearance: none;
	-webkit-appearance: none;
	width: 20px;
	height: 20px;
	border: 1px solid #fff;
	border-radius: 50%;
	margin-right: 10px;
	background-color: transparent;
	position: relative;
	top: 6px;
    transition: all .2s;
}
input[type="radio"]:checked {
    border-color: red;
    background-color: #00000024;
}

/* input[type="radio"]:checked::before {
	content: "";
	display: block;
	width: 12px;
	height: 12px;
	background-color: #ca8489;
	border-radius: 50%;
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
	animation: appear 0.2s;
} */
@keyframes appear {
	0% {
		transform: translate(-50%, -50%) scale(0);
	}
	100% {
		transform: translate(-50%, -50%) scale(1);
	}
}

</style>