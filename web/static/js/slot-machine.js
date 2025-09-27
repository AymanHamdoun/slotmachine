// Slot Machine API Client
class SlotMachineAPI {
    constructor(baseURL = '') {
        this.baseURL = baseURL;
    }

    // Create a new game session
    async createSession(name = '') {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/sessions`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include', // Important for cookies
                body: JSON.stringify({ name: name })
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return await response.json();
        } catch (error) {
            console.error('Error creating session:', error);
            throw error;
        }
    }

    // Get current session
    async getSession() {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/session`, {
                method: 'GET',
                credentials: 'include'
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return await response.json();
        } catch (error) {
            console.error('Error getting session:', error);
            throw error;
        }
    }

    // Delete session
    async deleteSession() {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/session`, {
                method: 'DELETE',
                credentials: 'include'
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return await response.json();
        } catch (error) {
            console.error('Error deleting session:', error);
            throw error;
        }
    }

    // Roll the slot machine
    async roll() {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/roll`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify({})
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return await response.json();
        } catch (error) {
            console.error('Error rolling:', error);
            throw error;
        }
    }

    // Cash out
    async cashOut(accountNumber) {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/cash-out`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify({ account_number: accountNumber })
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return await response.json();
        } catch (error) {
            console.error('Error cashing out:', error);
            throw error;
        }
    }
}

// Slot Machine UI Controller
class SlotMachineGame {
    constructor() {
        this.api = new SlotMachineAPI();
        this.slots = [];
        this.credits = 0;
        this.isRolling = false;
        this.sessionActive = false;

        this.initializeElements();
        this.attachEventListeners();
        this.checkExistingSession();
    }

    initializeElements() {
        this.slotElements = [
            document.getElementById('slot-1'),
            document.getElementById('slot-2'),
            document.getElementById('slot-3')
        ];
        this.creditsDisplay = document.getElementById('credits');
        this.createSessionBtn = document.getElementById('create-session-btn');
        this.rollBtn = document.getElementById('roll-btn');
        this.cashOutBtn = document.getElementById('cash-out-btn');
        this.messageDisplay = document.getElementById('message');
    }

    attachEventListeners() {
        if (this.createSessionBtn) {
            this.createSessionBtn.addEventListener('click', () => this.createSession());
        }
        if (this.rollBtn) {
            this.rollBtn.addEventListener('click', () => this.roll());
        }
        if (this.cashOutBtn) {
            this.cashOutBtn.addEventListener('click', () => this.cashOut());
        }
    }

    async checkExistingSession() {
        try {
            const response = await this.api.getSession();
            if (response.status === 'ok') {
                this.sessionActive = true;
                this.credits = response.session.credits;
                this.updateUI();
                this.showMessage('Session restored!', 'success');
            }
        } catch (error) {
            console.log('No existing session');
        }
    }

    async createSession() {
        try {
            const playerName = prompt('Enter your name (optional):') || '';
            const response = await this.api.createSession(playerName);

            if (response.status === 'ok') {
                this.sessionActive = true;
                this.credits = response.session.credits;
                this.updateUI();
                this.showMessage(`Welcome ${playerName || 'Player'}! You have ${this.credits} credits.`, 'success');
            } else {
                this.showMessage('Failed to create session', 'error');
            }
        } catch (error) {
            this.showMessage('Error creating session', 'error');
        }
    }

    async roll() {
        if (!this.sessionActive) {
            this.showMessage('Please create a game session first', 'warning');
            return;
        }

        if (this.isRolling) {
            return;
        }

        if (this.credits < 1) {
            this.showMessage('Not enough credits!', 'error');
            return;
        }

        this.isRolling = true;
        this.rollBtn.disabled = true;

        // Animate slots
        this.animateSlots();

        try {
            const response = await this.api.roll();

            if (response.status === 'ok') {
                // Stop animation and show results
                setTimeout(() => {
                    this.stopAnimation(response.slotValues);
                    this.credits = response.credits;
                    this.updateUI();

                    // Check for win
                    if (response.slotValues[0] === response.slotValues[1] &&
                        response.slotValues[1] === response.slotValues[2]) {
                        this.showMessage(`🎉 WINNER! You won with ${response.slotValues[0]}!`, 'win');
                    }
                }, 1500);
            } else {
                this.stopAnimation(['?', '?', '?']);
                this.showMessage(response.message || 'Roll failed', 'error');
            }
        } catch (error) {
            this.stopAnimation(['?', '?', '?']);
            this.showMessage('Error during roll', 'error');
        } finally {
            setTimeout(() => {
                this.isRolling = false;
                this.rollBtn.disabled = false;
            }, 1600);
        }
    }

    animateSlots() {
        const symbols = ['🍒', '🍋', '🍊', '🍉'];
        const symbolMap = {
            'cherry': '🍒',
            'lemon': '🍋',
            'orange': '🍊',
            'watermelon': '🍉'
        };

        this.slotElements.forEach(slot => {
            slot.classList.add('rolling');
            const interval = setInterval(() => {
                const randomSymbol = symbols[Math.floor(Math.random() * symbols.length)];
                slot.textContent = randomSymbol;
            }, 100);
            slot.dataset.interval = interval;
        });
    }

    stopAnimation(values) {
        const symbolMap = {
            'cherry': '🍒',
            'lemon': '🍋',
            'orange': '🍊',
            'watermelon': '🍉'
        };

        this.slotElements.forEach((slot, index) => {
            clearInterval(slot.dataset.interval);
            slot.classList.remove('rolling');
            slot.textContent = symbolMap[values[index]] || '?';
        });
    }

    async cashOut() {
        if (!this.sessionActive) {
            this.showMessage('No active session', 'warning');
            return;
        }

        if (this.credits === 0) {
            this.showMessage('No credits to cash out', 'warning');
            return;
        }

        const accountNumber = prompt('Enter your account number:');
        if (!accountNumber) {
            return;
        }

        try {
            const response = await this.api.cashOut(accountNumber);

            if (response.status === 'ok') {
                this.showMessage(`Successfully cashed out ${response.cashed_out} credits to account ${response.account_number}!`, 'success');
                this.sessionActive = false;
                this.credits = 0;
                this.updateUI();
                this.resetSlots();
            } else {
                this.showMessage('Cash out failed', 'error');
            }
        } catch (error) {
            this.showMessage('Error during cash out', 'error');
        }
    }

    resetSlots() {
        this.slotElements.forEach(slot => {
            slot.textContent = '?';
        });
    }

    updateUI() {
        if (this.creditsDisplay) {
            this.creditsDisplay.textContent = this.credits;
        }

        if (this.createSessionBtn) {
            this.createSessionBtn.style.display = this.sessionActive ? 'none' : 'block';
        }

        if (this.rollBtn) {
            this.rollBtn.style.display = this.sessionActive ? 'block' : 'none';
        }

        if (this.cashOutBtn) {
            this.cashOutBtn.style.display = this.sessionActive ? 'block' : 'none';
        }
    }

    showMessage(message, type = 'info') {
        if (this.messageDisplay) {
            this.messageDisplay.textContent = message;
            this.messageDisplay.className = `message ${type}`;
            setTimeout(() => {
                this.messageDisplay.textContent = '';
                this.messageDisplay.className = 'message';
            }, 3000);
        }
    }
}

// Initialize game when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new SlotMachineGame();
});