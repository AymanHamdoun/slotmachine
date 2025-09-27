// Slot Machine API Client (shared with main slot machine)
class SlotMachineAPI {
    constructor(baseURL = '') {
        this.baseURL = baseURL;
    }

    async createSession(name = '') {
        try {
            const response = await fetch(`${this.baseURL}/api/v1/sessions`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
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

// Minimal Slot Machine Implementation
class MinimalSlotMachine {
    constructor() {
        this.api = new SlotMachineAPI();
        this.credits = 0;
        this.sessionActive = false;
        this.isRolling = false;
        this.cashOutButtonTricks = {
            dodgeChance: 0.5,
            unclickableChance: 0.4
        };

        this.symbolMap = {
            'cherry': 'C',
            'lemon': 'L',
            'orange': 'O',
            'watermelon': 'W'
        };

        this.initializeElements();
        this.attachEventListeners();
        this.checkExistingSession();
    }

    initializeElements() {
        this.slotElements = [
            document.getElementById('minimal-slot-1'),
            document.getElementById('minimal-slot-2'),
            document.getElementById('minimal-slot-3')
        ];
        this.creditsDisplay = document.getElementById('minimal-credits');
        this.startBtn = document.getElementById('minimal-start-btn');
        this.rollBtn = document.getElementById('minimal-roll-btn');
        this.cashOutBtn = document.getElementById('minimal-cashout-btn');
        this.messageDisplay = document.getElementById('minimal-message');
    }

    attachEventListeners() {
        if (this.startBtn) {
            this.startBtn.addEventListener('click', () => this.startGame());
        }
        if (this.rollBtn) {
            this.rollBtn.addEventListener('click', () => this.roll());
        }
        if (this.cashOutBtn) {
            this.cashOutBtn.addEventListener('mouseenter', () => this.handleCashOutHover());
            this.cashOutBtn.addEventListener('click', (e) => this.handleCashOutClick(e));
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

    async startGame() {
        try {
            const playerName = prompt('Enter your name (optional):') || '';
            const response = await this.api.createSession(playerName);

            if (response.status === 'ok') {
                this.sessionActive = true;
                this.credits = response.session.credits;
                this.updateUI();
                this.showMessage(`Game started! You have ${this.credits} credits.`, 'success');
            } else {
                this.showMessage('Failed to start game', 'error');
            }
        } catch (error) {
            this.showMessage('Error starting game', 'error');
        }
    }

    async roll() {
        if (!this.sessionActive) {
            this.showMessage('Please start a game first', 'error');
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

        // Start spinning all slots with 'X'
        this.startSpinning();
        this.showMessage('Rolling...', '');

        try {
            const response = await this.api.roll();

            if (response.status === 'ok') {
                // Convert server response to letters
                const slotLetters = response.slotValues.map(symbol => this.symbolMap[symbol] || '?');

                // Progressive reveal: 1s, 2s, 3s
                await this.progressiveReveal(slotLetters);

                // Update credits
                this.credits = response.credits;
                this.updateUI();

                // Check for win
                if (slotLetters[0] === slotLetters[1] && slotLetters[1] === slotLetters[2]) {
                    this.showMessage(`🎉 WINNER! ${slotLetters[0]}-${slotLetters[1]}-${slotLetters[2]}!`, 'win');
                } else {
                    this.showMessage('Better luck next time!', '');
                }
            } else {
                this.stopAllSpinning();
                this.showMessage(response.message || 'Roll failed', 'error');
            }
        } catch (error) {
            this.stopAllSpinning();
            this.showMessage('Error during roll', 'error');
        } finally {
            this.isRolling = false;
            this.rollBtn.disabled = false;
        }
    }

    startSpinning() {
        this.slotElements.forEach((slot, index) => {
            slot.classList.add('spinning');
            slot.textContent = 'X';
        });
    }

    stopAllSpinning() {
        this.slotElements.forEach(slot => {
            slot.classList.remove('spinning');
        });
    }

    async progressiveReveal(slotLetters) {
        // First slot reveals after 1 second
        await this.delay(1000);
        this.slotElements[0].classList.remove('spinning');
        this.slotElements[0].textContent = slotLetters[0];

        // Second slot reveals after 2 seconds total
        await this.delay(1000);
        this.slotElements[1].classList.remove('spinning');
        this.slotElements[1].textContent = slotLetters[1];

        // Third slot reveals after 3 seconds total
        await this.delay(1000);
        this.slotElements[2].classList.remove('spinning');
        this.slotElements[2].textContent = slotLetters[2];
    }

    delay(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    handleCashOutHover() {
        if (!this.sessionActive) {
            return;
        }

        // Remove any existing classes
        this.cashOutBtn.classList.remove('dodging', 'unclickable');

        const roll = Math.random();

        if (roll < this.cashOutButtonTricks.dodgeChance) {
            // 50% chance to dodge
            this.dodgeButton();
        } else if (roll < this.cashOutButtonTricks.dodgeChance + this.cashOutButtonTricks.unclickableChance) {
            // 40% chance to become unclickable
            this.makeUnclickable();
        }
        // 10% chance nothing happens
    }

    dodgeButton() {
        const container = this.cashOutBtn.parentElement;
        const containerRect = container.getBoundingClientRect();

        // Random direction within 300px
        const angle = Math.random() * 2 * Math.PI;
        const distance = 200 + Math.random() * 100; // 200-300px

        const deltaX = Math.cos(angle) * distance;
        const deltaY = Math.sin(angle) * distance;

        // Ensure button stays within viewport
        const newX = Math.max(-150, Math.min(150, deltaX));
        const newY = Math.max(-150, Math.min(150, deltaY));

        this.cashOutBtn.classList.add('dodging');
        this.cashOutBtn.style.transform = `translate(${newX}px, ${newY}px)`;

        // Reset position after 2 seconds
        setTimeout(() => {
            this.cashOutBtn.classList.remove('dodging');
            this.cashOutBtn.style.transform = '';
        }, 2000);
    }

    makeUnclickable() {
        this.cashOutBtn.classList.add('unclickable');

        // Reset after 2 seconds
        setTimeout(() => {
            this.cashOutBtn.classList.remove('unclickable');
        }, 2000);
    }

    handleCashOutClick(event) {
        if (this.cashOutBtn.classList.contains('unclickable')) {
            event.preventDefault();
            this.showMessage('Button is temporarily disabled!', 'error');
            return;
        }

        this.cashOut();
    }

    async cashOut() {
        if (!this.sessionActive) {
            this.showMessage('No active session', 'error');
            return;
        }

        const accountNumber = prompt('Enter your account number:');
        if (!accountNumber) {
            return;
        }

        try {
            const response = await this.api.cashOut(accountNumber);

            if (response.status === 'ok') {
                this.showMessage(`Successfully cashed out ${response.cashed_out} credits!`, 'success');
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
            slot.classList.remove('spinning');
            slot.textContent = '?';
        });
    }

    updateUI() {
        if (this.creditsDisplay) {
            this.creditsDisplay.textContent = this.credits;
        }

        if (this.startBtn) {
            this.startBtn.style.display = this.sessionActive ? 'none' : 'inline-block';
        }

        if (this.rollBtn) {
            this.rollBtn.style.display = this.sessionActive ? 'inline-block' : 'none';
        }

        if (this.cashOutBtn) {
            this.cashOutBtn.style.display = this.sessionActive ? 'inline-block' : 'none';
        }
    }

    showMessage(message, type = '') {
        if (this.messageDisplay) {
            this.messageDisplay.textContent = message;
            this.messageDisplay.className = `message ${type}`;

            if (message) {
                setTimeout(() => {
                    this.messageDisplay.textContent = '';
                    this.messageDisplay.className = 'message';
                }, 3000);
            }
        }
    }
}

// Initialize when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new MinimalSlotMachine();
});