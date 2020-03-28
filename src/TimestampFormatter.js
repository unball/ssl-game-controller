const padWithZero = function (value) {
    if (value < 10) {
        return '0' + value.toString();
    }
    return value.toString();
};

const formatNsDuration = function (timestamp) {
    const seconds = Math.round(Math.abs(timestamp) / 1e9);
    const minutes = Math.floor(seconds / 60);
    const fullSeconds = seconds - minutes * 60;
    const sign = timestamp < 0 ? '-' : '';
    return sign + padWithZero(minutes) + ':' + padWithZero(fullSeconds);
};

const isNumeric = function (n) {
    return !isNaN(parseFloat(n)) && isFinite(n);
};

const process = function (el, binding) {
    let timestamp;
    if (isNumeric(binding.value)) {
        timestamp = binding.value;
    } else if (typeof binding.value === 'string' && binding.value.endsWith('s')) {
        let parts = binding.value.substring(0, binding.value.length - 1).split('.');
        if (parts.length === 1) {
            timestamp = parseInt(parts[0]) * 1000000000;
        } else if (parts.length === 2) {
            timestamp = parseInt(parts[0]) * 1000000000 + parseInt(parts[1]);
        } else {
            timestamp = el.innerHTML;
        }
    } else {
        timestamp = el.innerHTML;
    }
    if (!isNaN(timestamp)) {
        el.innerHTML = formatNsDuration(timestamp);
    }
};

const TimestampFormatter = {
    install(Vue) {
        Vue.directive('format-ns-duration', {
            bind(el, binding) {
                process(el, binding);
            },
            update(el, binding) {
                process(el, binding);
            }
        });
    }
};
export default TimestampFormatter;