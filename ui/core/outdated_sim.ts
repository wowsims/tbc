// Everything related to warning users that this sim is outdated and
// redirecting them to the new sim at wowsims.com.

const NEW_SIM_BASE_URL = 'https://www.wowsims.com/tbc/';

// Maps the old sim directory name (e.g. /tbc/hunter/) to the path of the
// matching sim on the new site. Specs which don't exist in the new sim
// (Smite Priest) link to the closest equivalent.
const newSimPaths: Record<string, string> = {
	'balance_druid': 'druid/balance/',
	'feral_druid': 'druid/feralcat/',
	'feral_tank_druid': 'druid/feralbear/',
	'elemental_shaman': 'shaman/elemental/',
	'enhancement_shaman': 'shaman/enhancement/',
	'hunter': 'hunter/dps/',
	'mage': 'mage/dps/',
	'rogue': 'rogue/dps/',
	'retribution_paladin': 'paladin/retribution/',
	'protection_paladin': 'paladin/protection/',
	'shadow_priest': 'priest/dps/',
	'smite_priest': 'priest/dps/',
	'warlock': 'warlock/dps/',
	'warrior': 'warrior/dps/',
	'protection_warrior': 'warrior/protection/',
	// The raid sim doesn't have an equivalent version of the new sim, so link to the landing page.
	'raid': '',
};

// Returns the URL of the matching sim on the new site.
export function getNewSimUrl(): string {
	const match = window.location.pathname.match(/\/tbc\/([a-z_]+)\/?/);
	const simDir = match ? match[1] : '';
	return NEW_SIM_BASE_URL + (newSimPaths[simDir] || '');
}

// Fixed bar at the very top of the page.
export function addOutdatedSimBar() {
	const newSimUrl = getNewSimUrl();
	const bar = document.createElement('div');
	bar.classList.add('outdated-sim-bar');
	bar.innerHTML = `
		<span>⚠️ This sim is outdated and no longer maintained! Use the new sim at <a href="${newSimUrl}">wowsims.com</a> instead. ⚠️</span>
	`;
	document.body.prepend(bar);
	document.body.classList.add('has-outdated-sim-bar');
}

// Warning text shown in the sidebar, below the iterations input.
export function makeOutdatedSimSidebarWarning(parentElem: HTMLElement) {
	const newSimUrl = getNewSimUrl();
	const isRaidSim = window.location.pathname.indexOf('/raid/') != -1;
	const exportNote = isRaidSim ? '' : `
		<br><br>
		You can bring your gear along:
		<ol>
			<li>Export &gt; New sim (top right)</li>
			<li>Import &gt; Addon in the new sim</li>
		</ol>
	`;
	const warning = document.createElement('div');
	warning.classList.add('outdated-sim-sidebar-warning');
	warning.innerHTML = `
		⚠️ This sim is <b>OUTDATED</b> and no longer maintained.
		<br><br>
		Use the new sim instead: <a href="${newSimUrl}">wowsims.com</a>
		${exportNote}
	`;
	parentElem.appendChild(warning);
}

// Modal shown on page load which redirects to the new sim after a few
// seconds, unless dismissed. Not shown when running locally (dev).
const REDIRECT_SECONDS = 10;
export function showOutdatedSimModal() {
	const hostname = window.location.hostname;
	const isLocal = hostname == 'localhost' || hostname == '127.0.0.1' || hostname == '0.0.0.0';
	const forceModal = window.location.search.includes('forceOutdatedModal');
	if (isLocal && !forceModal) {
		return;
	}

	const newSimUrl = getNewSimUrl();

	const overlay = document.createElement('div');
	overlay.classList.add('outdated-sim-modal-overlay');
	overlay.innerHTML = `
		<div class="outdated-sim-modal">
			<h2>⚠️ This sim is outdated! ⚠️</h2>
			<p>
				This version of the TBC sim was built for the original TBC Classic (2021) and is no longer maintained.
			</p>
			<p>
				A new, actively maintained version is available at <a href="${newSimUrl}">wowsims.com</a>.
			</p>
			<p class="outdated-sim-modal-countdown">
				Redirecting to the new sim in <span class="outdated-sim-modal-seconds">${REDIRECT_SECONDS}</span> seconds...
			</p>
			<div class="outdated-sim-modal-buttons">
				<button class="outdated-sim-modal-go">Go to the new sim now</button>
				<button class="outdated-sim-modal-dismiss">Dismiss (stay on the old sim)</button>
			</div>
		</div>
	`;
	document.body.appendChild(overlay);

	const secondsElem = overlay.getElementsByClassName('outdated-sim-modal-seconds')[0] as HTMLElement;
	let secondsLeft = REDIRECT_SECONDS;
	const countdownTimer = window.setInterval(() => {
		secondsLeft--;
		secondsElem.textContent = String(secondsLeft);
		if (secondsLeft <= 0) {
			window.clearInterval(countdownTimer);
			window.location.href = newSimUrl;
		}
	}, 1000);

	const goButton = overlay.getElementsByClassName('outdated-sim-modal-go')[0] as HTMLElement;
	goButton.addEventListener('click', () => {
		window.clearInterval(countdownTimer);
		window.location.href = newSimUrl;
	});

	const dismissButton = overlay.getElementsByClassName('outdated-sim-modal-dismiss')[0] as HTMLElement;
	dismissButton.addEventListener('click', () => {
		window.clearInterval(countdownTimer);
		overlay.remove();
	});
}
