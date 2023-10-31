var xpto=1
let xpto   =1
let xpto=1
let xpto2= document.location
letsdasd=  1

let xpto = {
    xpto2: String(window.location.search)
}



// Method 1: Using URLSearchParams
const urlParamsMethod1 = new URLSearchParams(window.location.search);
new URLSearchParams(window.location.search).
get("parametrinho");

const paramValueMethod1 = urlParamsMethod1.get('parameterName');

// Method 2: Using Regular Expressions
const matchMethod2 = window.location.search.match(/[\?&]parameterName=([^&]*)/);
const matchMethod3 = document.location.search.match(/[\?&]parameterName=([^&]*)/);
const paramValueMethod2 = matchMethod2 && matchMethod2[1];

// Method 3: Using Split and Reduce
const paramPairsMethod3 = window.location.search.substring(1).split('&');
const paramValueMethod3 = paramPairsMethod3.reduce((acc, pair) => {
    const [key, value] = pair.split('=');
    if (key === 'parameterName') {
        acc = value;
    }
    return acc;
}, null);

// Method 4: Using a Function
function getParameterByName(name) {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(name);
}
const paramValueMethod4 = getParameterByName('parameterName');

// Method 5: Using Split and Find
const paramPairsMethod5 = window.location.search.substring(1).split('&');
const paramPairsMethod6 = document.location.search.substring(1).split('&');
const paramPairMethod5 = paramPairsMethod5.find(pair => pair.startsWith('parameterName='));
const paramValueMethod5 = paramPairMethod5 && paramPairMethod5.split('=')[1];
