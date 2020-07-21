// Libraries
import React, {FunctionComponent} from 'react'

// Components
import {
  InputLabel,
  FlexBox,
  FlexDirection,
  ComponentSize,
  ComponentColor,
  SelectGroup,
  AlignItems,
  ButtonShape,
  Toggle,
  InputToggleType,
  QuestionMarkTooltip,
  ComponentStatus,
} from '@influxdata/clockface'
import SelectorList from 'src/timeMachine/components/SelectorList'
import BuilderCard from 'src/timeMachine/components/builderCard/BuilderCard'
import DurationInput from 'src/shared/components/DurationInput'

// Constants
import {DURATIONS} from 'src/timeMachine/constants/queryBuilder'

interface Props {
  isAutoWindowPeriod: boolean
  isAutoFunction: boolean
  isFillValues: boolean
  durationDisplay: string
  selectedFunctions: Array<string>
  functionList: Array<string>
  isInCheckOverlay: boolean
  onSelectFunction: (name: string) => void
  onChangeFillValues: () => void
  onSetFunctionSelectionMode: (m: 'custom' | 'auto') => void
  onSetWindowPeriodSelectionMode: (m: 'custom' | 'auto') => void
  onSelectAggregateWindow: (period: string) => void
}

const AggregationContents: FunctionComponent<Props> = ({
  isAutoWindowPeriod,
  isAutoFunction,
  isFillValues,
  durationDisplay,
  selectedFunctions,
  functionList,
  isInCheckOverlay,
  onSelectFunction,
  onChangeFillValues,
  onSetFunctionSelectionMode,
  onSetWindowPeriodSelectionMode,
  onSelectAggregateWindow,
}) => {
  const durationInputStatus = isAutoWindowPeriod
    ? ComponentStatus.Disabled
    : ComponentStatus.Default

  return (
    <BuilderCard className="aggregation-selector" testID="aggregation-selector">
      <BuilderCard.Header title="Window Period" />
      <BuilderCard.Body scrollable={false} style={{flex: 'unset'}}>
        <FlexBox
          direction={FlexDirection.Column}
          alignItems={AlignItems.Stretch}
          margin={ComponentSize.Small}
          stretchToFitWidth={true}
        >
          <SelectGroup shape={ButtonShape.StretchToFit}>
            <SelectGroup.Option
              name="custom"
              id="custom-window-period"
              active={!isAutoWindowPeriod}
              value="custom"
              onClick={onSetWindowPeriodSelectionMode}
              titleText="Custom"
            >
              Custom
            </SelectGroup.Option>
            <SelectGroup.Option
              name="auto"
              id="auto-window-period"
              active={isAutoWindowPeriod}
              value="auto"
              onClick={onSetWindowPeriodSelectionMode}
              titleText="Auto"
            >
              Auto
            </SelectGroup.Option>
          </SelectGroup>
          <DurationInput
            onSubmit={onSelectAggregateWindow}
            value={durationDisplay}
            suggestions={DURATIONS}
            submitInvalid={false}
            status={durationInputStatus}
          />
          <FlexBox
            direction={FlexDirection.Row}
            margin={ComponentSize.Small}
            stretchToFitWidth
            testID="component-spacer"
          >
            <Toggle
              id="isFillValues"
              type={InputToggleType.Checkbox}
              checked={isFillValues}
              onChange={onChangeFillValues}
              color={ComponentColor.Primary}
              size={ComponentSize.ExtraSmall}
            />
            <FlexBox.Child grow={1}>
              <InputLabel>Fill missing values</InputLabel>
            </FlexBox.Child>
            <QuestionMarkTooltip
              diameter={16}
              tooltipContents="Tooltip goes here!"
              tooltipStyle={{fontSize: '13px', padding: '8px'}}
            />
          </FlexBox>
        </FlexBox>
      </BuilderCard.Body>
      <BuilderCard.Header title="Aggregate Function" />
      <BuilderCard.Menu>
        <FlexBox
          direction={FlexDirection.Column}
          margin={ComponentSize.Small}
          alignItems={AlignItems.Stretch}
        >
          <SelectGroup shape={ButtonShape.StretchToFit}>
            <SelectGroup.Option
              name="custom"
              id="custom-function"
              active={!isAutoFunction}
              value="custom"
              onClick={onSetFunctionSelectionMode}
              titleText="Custom"
            >
              Custom
            </SelectGroup.Option>
            <SelectGroup.Option
              name="auto"
              id="auto-function"
              active={isAutoFunction}
              value="auto"
              onClick={onSetFunctionSelectionMode}
              titleText="Auto"
            >
              Auto
            </SelectGroup.Option>
          </SelectGroup>
        </FlexBox>
      </BuilderCard.Menu>
      <SelectorList
        items={functionList}
        selectedItems={selectedFunctions}
        onSelectItem={onSelectFunction}
        multiSelect={!isInCheckOverlay && !isAutoFunction}
      />
    </BuilderCard>
  )
}

export default AggregationContents
