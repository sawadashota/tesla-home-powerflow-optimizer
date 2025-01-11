<script lang="ts">
	import Textfield from "@smui/textfield";
	import Checkbox from '@smui/checkbox';
	import FormField from '@smui/form-field';
	import Button, { Label } from '@smui/button';
	import IconButton from '@smui/icon-button';
	import Snackbar, { Actions } from '@smui/snackbar';
	import { PUBLIC_API_URL } from '$env/static/public';

	export let data: ChargeSetting;
	let snackbar: Snackbar;

	async function submit(event: Event) {
		await fetch(`${PUBLIC_API_URL}/vehicle/charge/setting`, {
			method: "PUT",
			mode: "cors",
			credentials: "omit",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				enabled: data.enabled,
				charge_start_threshold: Number(data.charge_start_threshold),
				power_usage_increase_threshold: Number(data.power_usage_increase_threshold),
				power_usage_decrease_threshold: Number(data.power_usage_decrease_threshold),
				update_interval: Number(data.update_interval),
				minimum_setting: {
					threshold: Number(data.minimum_setting.threshold),
					time_range_start: data.minimum_setting.time_range_start,
					time_range_end: data.minimum_setting.time_range_end,
					amperage: Number(data.minimum_setting.amperage),
				},
			}),
		});
		snackbar.open();
	}
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<Snackbar bind:this={snackbar}>
		<Label>Update Successfully!</Label>
		<Actions>
			<IconButton class="material-icons" title="Dismiss">close</IconButton>
		</Actions>
	</Snackbar>
	<h1>
		Tesla Home PowerFlow Optimizer
	</h1>

	<h2>
		Setting
	</h2>

	<form on:submit|preventDefault={submit}>
		<div>
			<FormField>
				<Checkbox bind:checked={data.enabled} />
				<span slot="label">Enabled</span>
			</FormField>
		</div>
		<div>
			<Textfield
					style="width: 100%;"
					label="Charge Start Threshold (W)"
					bind:value={data.charge_start_threshold}
			/>
		</div>
		<div>
			<Textfield
					style="width: 100%;"
					label="Power Usage Increase Threshold (W)"
					bind:value={data.power_usage_increase_threshold}
			/>
		</div>
		<div>
			<Textfield
					style="width: 100%;"
					label="Power Usage Decrease Threshold (W)"
					bind:value={data.power_usage_decrease_threshold}
			/>
		</div>
		<div>
			<Textfield
					style="width: 100%;"
					label="Update Interval (minutes)"
					bind:value={data.update_interval}
			/>
		</div>
		<div>
			<Textfield
					style="width: 100%;"
					label="Minimum Charge Level (%)"
					bind:value={data.minimum_setting.threshold}
			/>
		</div>
		<div>
			<Textfield
					style="width: 100%;"
					label="Minimum Charge Time Range Start (HH:MM)"
					bind:value={data.minimum_setting.time_range_start}
			/>
		</div>
		<div>
			<Textfield
					style="width: 100%;"
					label="Minimum Charge Time Range End (HH:MM)"
					bind:value={data.minimum_setting.time_range_end}
			/>
		</div>
		<div>
			<Textfield
					style="width: 100%;"
					label="Amperage (A) for Minimum Charge"
					bind:value={data.minimum_setting.amperage}
			/>
		</div>

		<div class="field-submit">
			<Button color="secondary" variant="outlined">
				<Label>Save</Label>
			</Button>
		</div>
	</form>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
		font-size: 1.6em;
	}

	h2 {
		font-size: 1.4em;
	}

	form {
		width: 80%;
		max-width: 300px;
	}

	.field-submit {
		margin-top: 2em;
		text-align: center;
	}
</style>
