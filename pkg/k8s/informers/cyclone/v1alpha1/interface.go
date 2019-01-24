/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Projects returns a ProjectInformer.
	Projects() ProjectInformer
	// Resources returns a ResourceInformer.
	Resources() ResourceInformer
	// Stages returns a StageInformer.
	Stages() StageInformer
	// Workflows returns a WorkflowInformer.
	Workflows() WorkflowInformer
	// WorkflowRuns returns a WorkflowRunInformer.
	WorkflowRuns() WorkflowRunInformer
	// WorkflowTriggers returns a WorkflowTriggerInformer.
	WorkflowTriggers() WorkflowTriggerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Projects returns a ProjectInformer.
func (v *version) Projects() ProjectInformer {
	return &projectInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Resources returns a ResourceInformer.
func (v *version) Resources() ResourceInformer {
	return &resourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Stages returns a StageInformer.
func (v *version) Stages() StageInformer {
	return &stageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Workflows returns a WorkflowInformer.
func (v *version) Workflows() WorkflowInformer {
	return &workflowInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WorkflowRuns returns a WorkflowRunInformer.
func (v *version) WorkflowRuns() WorkflowRunInformer {
	return &workflowRunInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WorkflowTriggers returns a WorkflowTriggerInformer.
func (v *version) WorkflowTriggers() WorkflowTriggerInformer {
	return &workflowTriggerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
