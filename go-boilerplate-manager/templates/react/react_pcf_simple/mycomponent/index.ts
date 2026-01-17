import { IInputs, IOutputs } from "./generated/ManifestTypes";

export class MyComponent implements ComponentFramework.StandardControl<IInputs, IOutputs> {
	// Current value of the control
	private _value: string;

	// Callback function to notify framework of changes
	private _notifyOutputChanged: () => void;

	// Input element reference
	private _inputElement: HTMLInputElement;

	// Container element reference
	private _container: HTMLDivElement;

	/**
	 * Empty constructor.
	 */
	constructor() {
		// No initialization needed here
	}

	/**
	 * Initialize the control instance.
	 * @param context The context object containing properties and utility functions
	 * @param notifyOutputChanged Callback to notify framework of property changes
	 * @param state Persisted state from previous session
	 * @param container The container div element to render the control
	 */
	public init(
		context: ComponentFramework.Context<IInputs>,
		notifyOutputChanged: () => void,
		state: ComponentFramework.Dictionary,
		container: HTMLDivElement
	): void {
		// Store the callback
		this._notifyOutputChanged = notifyOutputChanged;

		// Create main container
		this._container = document.createElement("div");
		this._container.className = "mycomponent-container";

		// Create input element
		this._inputElement = document.createElement("input");
		this._inputElement.type = "text";
		this._inputElement.className = "mycomponent-input";

		// Add event listeners
		this._inputElement.addEventListener("input", this.onInputChange.bind(this));
		this._inputElement.addEventListener("blur", this.onInputBlur.bind(this));

		// Append input to container
		this._container.appendChild(this._inputElement);

		// Append container to the control container
		container.appendChild(this._container);
	}

	/**
	 * Handle input change event
	 */
	private onInputChange(event: Event): void {
		this._value = this._inputElement.value;
		this._notifyOutputChanged();
	}

	/**
	 * Handle input blur event
	 */
	private onInputBlur(event: Event): void {
		this._value = this._inputElement.value;
		this._notifyOutputChanged();
	}

	/**
	 * Update the control view with new data
	 * @param context The context object containing updated properties
	 */
	public updateView(context: ComponentFramework.Context<IInputs>): void {
		// Update the value from context
		this._value = context.parameters.Value.raw || "";
		this._inputElement.value = this._value;

		// Update placeholder
		const placeholder = context.parameters.Placeholder.raw;
		if (placeholder) {
			this._inputElement.placeholder = placeholder;
		}

		// Update readonly state
		const readOnly = context.parameters.ReadOnly.raw;
		this._inputElement.readOnly = readOnly || false;

		// Handle validation errors
		if (context.parameters.Value.error) {
			this._inputElement.classList.add("mycomponent-input-error");
		} else {
			this._inputElement.classList.remove("mycomponent-input-error");
		}
	}

	/**
	 * Return outputs to the framework
	 * @returns An object containing the output properties
	 */
	public getOutputs(): IOutputs {
		return {
			Value: this._value
		};
	}

	/**
	 * Cleanup when the control is destroyed
	 */
	public destroy(): void {
		// Remove event listeners
		if (this._inputElement) {
			this._inputElement.removeEventListener("input", this.onInputChange.bind(this));
			this._inputElement.removeEventListener("blur", this.onInputBlur.bind(this));
		}
	}
}
