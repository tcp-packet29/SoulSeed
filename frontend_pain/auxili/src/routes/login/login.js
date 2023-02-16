/**
 * @param {string} item
 */
function setItem(item) {
    window.localStorage.setItem('token', item);
}

function getItem() {
    return window.localStorage.getItem('token');
}

export default {
    setItem,
    getItem,
}