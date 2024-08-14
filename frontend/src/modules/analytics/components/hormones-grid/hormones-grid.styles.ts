import { cva, type VariantProps } from 'class-variance-authority'

export const hormonesVariants = cva('', {
	variants: {
		variant: {
			endorphin: 'bg-[#E55733]',
			dopamine: 'bg-[#B286FE]',
			serotonin: 'bg-[#FD9894]',
			oxytocin: 'bg-[#B2F042]',
		},
	},
	defaultVariants: {
		variant: 'endorphin',
	},
})

export type HormonesVariants = VariantProps<typeof hormonesVariants>
