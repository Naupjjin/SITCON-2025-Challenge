document.getElementById('addEnv').onclick = function() {
    const envContainer = document.getElementById('envVariables');
    const newEnvVariable = document.createElement('div');
    newEnvVariable.className = 'env-variable';
    newEnvVariable.innerHTML = `
        <label for="key">ENV name:</label>
        <input type="text" class="key" placeholder="input your ENV name">
        <label for="value">ENV Value:</label>
        <input type="text" class="value" placeholder="input your ENV value">
    `;
    envContainer.appendChild(newEnvVariable);
};

document.getElementById('submit').onclick = async function() {
    const envDict = {};
    const keys = document.querySelectorAll('.key');
    const values = document.querySelectorAll('.value');
    
    for (let i = 0; i < keys.length; i++) {
        if (keys[i].value && values[i].value) {
            envDict[keys[i].value] = values[i].value;
        }
    }

    const code = document.getElementById('code').value;

    const response = await fetch('/mygolang', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ env: envDict, code: code }),
    });
    if (!response.ok) {
        const errorMessage = await response.text();
        alert("Error: " + errorMessage); 
    } else {
        alert("MyGo!!!!!"); 
    }
};
