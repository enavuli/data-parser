const fs = require('fs');
const path = require('path');

class Parser {
    constructor(filePath) {
        this.filePath = path.resolve(filePath);
    }

    parse() {
        if (!fs.existsSync(this.filePath)) {
            throw new Error(`File not found: ${this.filePath}`);
        }

        const fileContent = fs.readFileSync(this.filePath, 'utf-8');
        const lines = fileContent.split('\n');
        const parsedData = lines.map(line => {
            const [key, value] = line.split('=');
            return { key: key.trim(), value: value ? value.trim() : null };
        });

        return parsedData.filter(entry => entry.key);
    }

    static validate(data) {
        return data.every(entry => entry.key && entry.value);
    }
}

module.exports = Parser;